package rules

import "github.com/tiwarishub/restaurant-recomm/src/models"

type Rules interface {
	ApplyRule(u models.User, r []models.Restaurant) ([]models.Restaurant, error)
}
