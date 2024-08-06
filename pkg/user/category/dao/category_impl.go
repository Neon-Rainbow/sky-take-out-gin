package dao

import (
	"context"
	"gorm.io/gorm"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type CategoryDaoImpl struct {
	db database.DatabaseInterface
}

// GetCategoryList 获取分类列表
func (dao *CategoryDaoImpl) GetCategoryList(ctx context.Context, Type int) ([]model.Category, error) {
	var categories []model.Category
	err := dao.db.GetDB().WithContext(ctx).Where("type = ?", Type).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryWithTransaction 获取分类列表
func (dao *CategoryDaoImpl) GetCategoryWithTransaction(ctx context.Context, tx *gorm.DB, Type int) ([]model.Category, error) {
	var categories []model.Category
	err := tx.WithContext(ctx).Where("type = ?", Type).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// BeginTransaction 开启事务
func (dao *CategoryDaoImpl) BeginTransaction() *gorm.DB {
	return dao.db.GetDB().Begin()
}

func NewCategoryDao(db database.DatabaseInterface) *CategoryDaoImpl {
	return &CategoryDaoImpl{db}
}
