package models

import "time"

type Restaurant struct {
	ID            string
	Cuisine       Cuisine
	CostBracket   int
	Rating        float32
	IsRecommended bool
	OnBoardedTime time.Time
}
