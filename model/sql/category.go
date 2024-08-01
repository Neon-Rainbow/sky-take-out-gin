package model

import "time"

type Category struct {
	ID         int64     `json:"id" gorm:"primary_key"`
	Type       int       `json:"type"`
	Name       string    `json:"name" binding:"required"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateUser int64     `json:"create_user"`
	UpdateUser int64     `json:"update_user"`
}