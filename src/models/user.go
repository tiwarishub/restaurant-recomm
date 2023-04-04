package models

type User struct {
	ID          int
	Cuisines    []CuisineTracking
	CostBracket []CostTracking
}

type CostTracking struct {
	Cost       int
	NoOfOrders uint64
}

type CuisineTracking struct {
	Cuisine    Cuisine
	NoOfOrders uint64
}
