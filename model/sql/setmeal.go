package model

import "time"

type Setmeal struct {
	ID            int64         `json:"id" gorm:"primary_key;autoIncrement"`
	CategoryID    int64         `json:"category_id"`
	Name          string        `json:"name"`
	Price         float64       `json:"price"`
	Status        int           `json:"status"`
	Description   string        `json:"description"`
	Image         string        `json:"image"`
	CreateTime    time.Time     `json:"create_time"`
	UpdateTime    time.Time     `json:"update_time"`
	CreateUser    int64         `json:"create_user"`
	UpdateUser    int64         `json:"update_user"`
	SetmealDishes []SetmealDish `json:"setmeal_dishes" gorm:"foreignKey:SetmealID;references:ID"`
}

func (Setmeal) TableName() string {
	return "setmeal"
}
