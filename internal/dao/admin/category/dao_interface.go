package category

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type CategoryDao interface {
	UpdateCategoryType(ctx context.Context, category *model.Category) error
	GetCategoryById(ctx context.Context, id int64) (*model.Category, error)
	GetCategoryPage(ctx context.Context, name string, page, pageSize, categoryType int) (categories []model.Category, total int64, err error)
	ChangeCategoryStatus(ctx context.Context, id int64, status int) error
}
