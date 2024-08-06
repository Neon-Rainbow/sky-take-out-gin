package model

import "gorm.io/gorm"

type SetmealDish struct {
	gorm.Model
	SetmealID uint    `gorm:"comment:'套餐id'"`
	DishID    uint    `gorm:"comment:‘菜品id’"`
	Name      string  `gorm:"type:varchar(32);comment:'菜品名称(冗余字段)'"`
	Price     float64 `gorm:"type:decimal(10,2);comment:‘菜品单价（冗余字段）’"`
	Copies    int     `gorm:"comment:‘菜品份数’"`
}
