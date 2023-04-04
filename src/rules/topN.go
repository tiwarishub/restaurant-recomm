package rules

import (
	"sort"

	"github.com/tiwarishub/restaurant-recomm/src/models"
)

type TopN struct {
	Top int
}

func (t TopN) ApplyRule(u models.User, r []models.Restaurant) ([]models.Restaurant, error) {
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].OnBoardedTime.Compare(r[j].OnBoardedTime) == 1
	})
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Rating > r[j].Rating
	})
	var limit = t.Top
	if len(r) < limit {
		limit = len(r)
	}

	return r[:limit], nil

}
