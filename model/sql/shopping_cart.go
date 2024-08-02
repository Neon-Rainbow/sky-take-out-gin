package model

import "time"

type ShoppingCart struct {
	ID         int64     `json:"id" gorm:"primary_key;autoIncrement"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	UserID     int64     `json:"user_id"`
	DishID     int64     `json:"dish_id"`
	SetmealID  int64     `json:"setmeal_id"`
	DishFlavor string    `json:"dish_flavor"`
	Number     int       `json:"number"`
	Amount     float64   `json:"amount"`
	CreateTime time.Time `json:"create_time"`
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
