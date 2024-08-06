package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type DishFlavor struct {
	gorm.Model
	DishID uint            `gorm:"not null;comment:'菜品'"`
	Name   string          `gorm:"type:varchar(32);comment:'口味名称'"`
	Value  json.RawMessage `gorm:"type:json;comment:'口味数据list'"`
}
