package model

// DishFlavor 菜品口味
type DishFlavor struct {
	ID     int64  `json:"id" gorm:"primary_key;autoIncrement"`
	DishID int64  `json:"dish_id"`
	Name   string `json:"name" gorm:"type:longtext"`
	Value  string `json:"value"`
}

func (DishFlavor) TableName() string {
	return "dish_flavor"
}
