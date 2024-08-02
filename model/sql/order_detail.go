package model

type OrderDetail struct {
	ID         int64   `json:"id" gorm:"primary_key;autoIncrement"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	OrderID    int64   `json:"order_id"`
	DishID     int64   `json:"dish_id"`
	SetmealID  int64   `json:"setmeal_id"`
	DishFlavor string  `json:"dish_flavor"`
	Number     int     `json:"number"`
	Amount     float64 `json:"amount"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
