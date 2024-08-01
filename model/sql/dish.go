package model

import "time"

type Dish struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" binding:"required"`
	CategoryID  int64     `json:"category_id" binding:"required"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	CreateUser  int64     `json:"create_user"`
	UpdateUser  int64     `json:"update_user"`
}
