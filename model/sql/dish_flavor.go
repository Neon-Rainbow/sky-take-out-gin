package model

// DishFlavor 菜品口味
type DishFlavor struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	DishID int64  `json:"dish_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}
