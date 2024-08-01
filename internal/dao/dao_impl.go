package dao

import (
	"gorm.io/gorm"
	"sky-take-out-gin/utils/MySQL"
)

type CategoryDaoImpl struct {
	*gorm.DB
}

// NewCategoryDaoImpl 实例化CategoryDaoImpl
func NewCategoryDaoImpl() *CategoryDaoImpl {
	db := MySQL.GetDB()
	return &CategoryDaoImpl{
		db,
	}
}
