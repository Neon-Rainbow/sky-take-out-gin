package model

import (
	"gorm.io/gorm"
)

type Setmeal struct {
	gorm.Model
	CategoryID  uint    `gorm:"not null;comment:'菜品分类id'"`
	Name        string  `gorm:"type:varchar(32);not null;unique;comment:'套餐名称'"`
	Price       float64 `gorm:"type:decimal(10,2);not null;comment:'套餐价格'"`
	Status      int     `gorm:"default:1;comment:'售卖状态 0:停售 1:起售'"`
	Description string  `gorm:"type:varchar(255);comment:'描述信息'"`
	Image       string  `gorm:"type:varchar(255);comment:'图片'"`
	CreateUser  uint    `gorm:"comment:'创建人'"`
	UpdateUser  uint    `gorm:"comment:'修改人'"`
}
