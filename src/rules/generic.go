package rules

import "github.com/tiwarishub/restaurant-recomm/src/models"

type Generic struct {
	RatingCheck      func(float32) bool
	CuisineCheck     func(userCusines []models.CuisineTracking, resCuisine models.Cuisine) bool
	CostBracketCheck func(userCostBracker []models.CostTracking, resCost int) bool
	IsFeaturedCheck  bool
	Cuisine          []models.CuisineTracking
	CostBracket      []models.CostTracking
}

func (g Generic) ApplyRule(u models.User, r []models.Restaurant) ([]models.Restaurant, error) {
	var filteredRestaurants []models.Restaurant
	for _, restaurant := range r {
		pass := true
		if g.IsFeaturedCheck && !restaurant.IsRecommended {
			pass = false
		}

		if g.CuisineCheck != nil && !g.CuisineCheck(u.Cuisines, restaurant.Cuisine) {
			pass = false
		}

		if g.CostBracketCheck != nil && !g.CostBracketCheck(u.CostBracket, restaurant.CostBracket) {
			pass = false
		}

		if g.RatingCheck != nil && !g.RatingCheck(restaurant.Rating) {
			pass = false
		}

		if pass {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}

	return filteredRestaurants, nil
}
