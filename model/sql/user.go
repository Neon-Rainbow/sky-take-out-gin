package model

import "time"

type User struct {
	ID         int64     `json:"id" gorm:"primary_key"`
	OpenID     string    `json:"openid"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IDNumber   string    `json:"id_number"`
	Avatar     string    `json:"avatar"`
	CreateTime time.Time `json:"create_time"`
}

func (User) TableName() string {
	return "user"
}
