package model

type SetmealDish struct {
	ID        int64   `json:"id" gorm:"primary_key"`
	SetmealID int64   `json:"setmeal_id"`
	DishID    int64   `json:"dish_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Copies    int     `json:"copies"`
}