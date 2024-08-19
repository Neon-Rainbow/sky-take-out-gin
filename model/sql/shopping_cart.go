package model

import (
	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model `copier:"-"`
	Name       string  `gorm:"type:varchar(32);comment:'商品名称'"`
	Image      string  `gorm:"type:varchar(255);comment:'图片'"`
	UserID     uint    `gorm:"not null;comment:'主键'"`
	DishID     uint    `gorm:"comment:'菜品id'"`
	SetmealID  uint    `gorm:"comment:'套餐id'"`
	DishFlavor string  `gorm:"type:varchar(50);comment:'口味'"`
	Number     int     `gorm:"default:1;comment:'数量'"`
	Amount     float64 `gorm:"type:decimal(10,2);not null;comment:'金额'"`
}
