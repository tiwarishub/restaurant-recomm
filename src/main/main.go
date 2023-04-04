package main

import (
	"fmt"
	"log"

	"github.com/tiwarishub/restaurant-recomm/src/loader"
	"github.com/tiwarishub/restaurant-recomm/src/models"
	"github.com/tiwarishub/restaurant-recomm/src/recommender"
)

func main() {
	restaurants, err := loader.GetRestaurantsData()
	if err != nil {
		log.Fatal("Unable to load the data. Exiting", err)
	}

	// create a sample user
	user := createSampleUser()

	// getRestaurantRecommendations
	recommendedList, err := getRestaurantRecommendations(user, restaurants)
	if err != nil {
		log.Fatal("Unable to get resturants recommendations:", err)
	}

	// print the recommended list
	for _, id := range recommendedList {
		fmt.Println(id)
	}
}

func getRestaurantRecommendations(user models.User, restaurants []models.Restaurant) ([]string, error) {
	return recommender.GetRestaurants(user, restaurants)
}

func createSampleUser() models.User {
	return models.User{
		ID: 1,
		Cuisines: []models.CuisineTracking{
			{
				Cuisine:    models.NorthIndiaCuisine,
				NoOfOrders: 100,
			},
			{
				Cuisine:    models.ChineseCuisine,
				NoOfOrders: 50,
			},
		},
		CostBracket: []models.CostTracking{
			{
				Cost:       5,
				NoOfOrders: 50,
			},
			{
				Cost:       2,
				NoOfOrders: 100,
			},
		},
	}
}
