package dao

import (
	"sky-take-out-gin/pkg/common/database"
)

type CategoryDaoImpl struct {
	db database.DatabaseInterface
}

// NewCategoryDaoImpl 实例化CategoryDaoImpl
func NewCategoryDaoImpl(db database.DatabaseInterface) *CategoryDaoImpl {
	return &CategoryDaoImpl{
		db,
	}
}
