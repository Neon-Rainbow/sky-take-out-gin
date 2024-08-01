package model

import "time"

type Employee struct {
	ID         int64     `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IDNumber   string    `json:"id_number"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateUser int64     `json:"create_user"`
	UpdateUser int64     `json:"update_user"`
}

func (Employee) TableName() string {
	return "employee"
}
