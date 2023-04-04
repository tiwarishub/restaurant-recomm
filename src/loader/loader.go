package loader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tiwarishub/restaurant-recomm/src/models"
)

// const dataFilePath = "./data.csv"
const dataFilePath = "/Users/shubham/go/src/github.com/tiwarishub/restaurant-recomm/data.csv"

func GetRestaurantsData() ([]models.Restaurant, error) {

	fd, err := os.Open(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open a file: %w", err)
	}

	defer fd.Close()
	reader := csv.NewReader(fd)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read source file: %w", err)
	}

	return parseFileData(data)
}

func parseFileData(data [][]string) ([]models.Restaurant, error) {

	var restaurants []models.Restaurant
	for _, line := range data {
		restaurant, _ := parseLine(line)
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func parseLine(fields []string) (models.Restaurant, error) {
	var restaurant models.Restaurant
	var err error
	for index, field := range fields {
		switch index {
		case 0:
			restaurant.ID = field
		case 1:
			cuisineInt, err := strconv.Atoi(field)
			if err != nil {
				restaurant.Cuisine = models.UnknownCuisine
				break
			}

			restaurant.Cuisine = models.Cuisine(cuisineInt)
		case 2:
			restaurant.CostBracket, err = strconv.Atoi(field)
			if err != nil {
				restaurant.CostBracket = -1
			}
		case 3:
			rating, err := strconv.ParseFloat(field, 32)
			if err != nil {
				restaurant.Rating = -1
			}
			restaurant.Rating = float32(rating)
		case 4:
			restaurant.IsRecommended, err = strconv.ParseBool(field)
			if err != nil {
				restaurant.IsRecommended = false
			}
		case 5:
			restaurant.OnBoardedTime, err = time.Parse("2023-01-31", field)
			if err != nil {
				restaurant.OnBoardedTime = time.Now()
			}
		}
	}

	return restaurant, nil
}
