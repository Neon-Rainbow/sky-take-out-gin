package model

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model  `copier:"-"`
	Name        string       `gorm:"type:varchar(32);not null;unique;comment:'菜品名称'"`
	CategoryID  uint         `gorm:"not null;comment:'菜品分类id'"`
	Price       float64      `gorm:"type:decimal(10,2);comment:'菜品价格'"`
	Image       string       `gorm:"type:varchar(255);comment:'图片'"`
	Description string       `gorm:"type:varchar(255);comment:'描述信息'"`
	Status      int          `gorm:"default:1;comment:'0 停售 1 起售'"`
	CreateUser  uint         `gorm:"comment:'创建人'"`
	UpdateUser  uint         `gorm:"comment:'修改人'"`
	DishFlavors []DishFlavor `gorm:"foreignKey:DishID"`
}
