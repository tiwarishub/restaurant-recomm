package recommender

import (
	"github.com/tiwarishub/restaurant-recomm/src/models"
	"github.com/tiwarishub/restaurant-recomm/src/rules"
)

func ApplyRules(user models.User, restaurants []models.Restaurant, rules []rules.Rules) ([]models.Restaurant, error) {
	var r []models.Restaurant
	for _, rule := range rules {
		ruleRecommendation, _ := rule.ApplyRule(user, restaurants)
		r = append(r, ruleRecommendation...)
		restaurants = difference(restaurants, r)
	}

	return r, nil
}

func PrimaryCusineCheck(uc []models.CuisineTracking, rc models.Cuisine) bool {
	return rc == uc[0].Cuisine
}

func SecondaryCusineCheck(ucs []models.CuisineTracking, rc models.Cuisine) bool {
	for _, uc := range ucs[1:] {
		if uc.Cuisine == rc {
			return true
		}
	}

	return false
}

func PrimaryCostBracketChecks(uc []models.CostTracking, rc int) bool {
	return rc == uc[0].Cost
}

func SecondaryCostBracketChecks(ucs []models.CostTracking, rc int) bool {
	for _, uc := range ucs[1:] {
		if uc.Cost == rc {
			return true
		}
	}

	return false
}

// a-b
func difference(aRestaurants []models.Restaurant, bRestaurants []models.Restaurant) []models.Restaurant {
	var result []models.Restaurant
	bLen := len(bRestaurants)
	if bLen == 0 {
		return aRestaurants
	}

	bMap := make(map[string]bool, bLen)
	for _, r := range bRestaurants {
		bMap[r.ID] = true
	}

	for _, r := range aRestaurants {
		if _, ok := bMap[r.ID]; !ok {
			result = append(result, r)
		}
	}

	return result
}
