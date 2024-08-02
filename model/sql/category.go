package model

import (
	"time"
)

type Category struct {
	ID         int64     `json:"id" gorm:"primary_key;autoIncrement"`
	Type       int       `json:"type"`
	Name       string    `json:"name" binding:"required" gorm:"type:longtext"` // 指定索引前缀长度
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateUser int64     `json:"create_user"`
	UpdateUser int64     `json:"update_user"`
}

// TableName 指定表名为 custom_category
func (Category) TableName() string {
	return "category"
}
