package model

import (
	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model `copier:"-"`
	Name       string  `gorm:"type:varchar(32);comment:'名字'"`
	Image      string  `gorm:"type:varchar(255);comment:'图片'"`
	OrderID    uint    `gorm:"not null;comment:'订单id'"`
	DishID     uint    `gorm:"comment:'菜品id'"`
	SetmealID  uint    `gorm:"comment:'套餐id'"`
	DishFlavor string  `gorm:"type:varchar(50);comment:'口味'"`
	Number     int     `gorm:"default:1;comment:'数量'"`
	Amount     float64 `gorm:"type:decimal(10,2);not null;comment:'金额'"`
}
