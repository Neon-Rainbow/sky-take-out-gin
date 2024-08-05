package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type CategoryDaoInfertace interface {
	UpdateCategoryType(ctx context.Context, category *model.Category) error
	GetCategoryById(ctx context.Context, id int64) (*model.Category, error)
	GetCategoryPage(ctx context.Context, name string, page, pageSize, categoryType int) (categories []model.Category, total int64, err error)
	ChangeCategoryStatus(ctx context.Context, id int64, status int) error
	CreateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, id int64) error
	GetCategoryByType(ctx context.Context, categoryType int) ([]model.Category, error)
}