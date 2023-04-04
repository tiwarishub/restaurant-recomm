package recommender

import (
	"github.com/tiwarishub/restaurant-recomm/src/models"
	"github.com/tiwarishub/restaurant-recomm/src/rules"
)

func GetRestaurants(user models.User, restaurants []models.Restaurant) ([]string, error) {
	var recommededRestaurant []models.Restaurant

	// Create rules
	rule1a := &rules.Generic{CuisineCheck: PrimaryCusineCheck, CostBracketCheck: PrimaryCostBracketChecks, IsFeaturedCheck: true}
	rule1b := &rules.Generic{CuisineCheck: PrimaryCusineCheck, CostBracketCheck: SecondaryCostBracketChecks, IsFeaturedCheck: true}
	rule1c := &rules.Generic{CuisineCheck: SecondaryCusineCheck, CostBracketCheck: PrimaryCostBracketChecks, IsFeaturedCheck: true}
	rule2 := &rules.Generic{CuisineCheck: PrimaryCusineCheck, CostBracketCheck: PrimaryCostBracketChecks, RatingCheck: func(rating float32) bool { return rating >= 4 }}
	rule3 := &rules.Generic{CuisineCheck: PrimaryCusineCheck, CostBracketCheck: SecondaryCostBracketChecks, RatingCheck: func(rating float32) bool { return rating >= 4.5 }}
	rule4 := &rules.Generic{CuisineCheck: SecondaryCusineCheck, CostBracketCheck: PrimaryCostBracketChecks, RatingCheck: func(rating float32) bool { return rating >= 4.5 }}
	rule5 := &rules.TopN{Top: 4}
	rule6 := &rules.Generic{CuisineCheck: PrimaryCusineCheck, CostBracketCheck: PrimaryCostBracketChecks, RatingCheck: func(rating float32) bool { return rating < 4 }}
	rule7 := &rules.Generic{CuisineCheck: PrimaryCusineCheck, CostBracketCheck: SecondaryCostBracketChecks, RatingCheck: func(rating float32) bool { return rating < 4.5 }}
	rule8 := &rules.Generic{CuisineCheck: SecondaryCusineCheck, CostBracketCheck: PrimaryCostBracketChecks, RatingCheck: func(rating float32) bool { return rating < 4.5 }}

	// Applying rules

	//Rules 1a,b and c
	rule1Recommendation, _ := rule1a.ApplyRule(user, restaurants)

	if len(rule1Recommendation) == 0 {
		rule1Recommendation, _ = ApplyRules(user, restaurants, []rules.Rules{rule1b, rule1c})
	}

	// Applying remaining rules
	recommededRestaurant, _ = ApplyRules(user, difference(restaurants, rule1Recommendation), []rules.Rules{rule2, rule3, rule4, rule5, rule6, rule7, rule8})

	// final list of restaurants
	recommededRestaurant = append(rule1Recommendation, recommededRestaurant...)
	unOrderedRestaurant := difference(restaurants, recommededRestaurant)

	recommededRestaurant = append(recommededRestaurant, unOrderedRestaurant...)
	var restaurantIDs []string
	for _, r := range recommededRestaurant {
		restaurantIDs = append(restaurantIDs, r.ID)
	}

	var limit int
	if len(restaurantIDs) > 100 {
		limit = 100
	} else {
		limit = len(restaurantIDs)
	}
	return restaurantIDs[:limit], nil
}
