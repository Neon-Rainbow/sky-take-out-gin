package model

import (
	"gorm.io/gorm"
)

// Category 分类
type Category struct {
	gorm.Model
	Type       int    `gorm:"comment:'类型 1 菜品分类 2 套餐分类'"`
	Name       string `gorm:"type:varchar(32);not null;unique;comment:'分类名称'"`
	Sort       int    `gorm:"default:0;comment:'顺序'"`
	Status     int    `gorm:"comment:'分类状态 0:禁用，1:启用'"`
	CreateUser uint   `gorm:"comment:'创建人'"`
	UpdateUser uint   `gorm:"comment:'修改人'"`
}
