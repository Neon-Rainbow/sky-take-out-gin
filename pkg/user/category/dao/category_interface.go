package dao

import (
	"context"
	"gorm.io/gorm"
	model "sky-take-out-gin/model/sql"
)

type CategoryDaoInterface interface {
	// GetCategoryList 获取分类列表
	GetCategoryList(ctx context.Context, Type int) ([]model.Category, error)
	GetCategoryWithTransaction(ctx context.Context, tx *gorm.DB, Type int) ([]model.Category, error)
	BeginTransaction() *gorm.DB
}
