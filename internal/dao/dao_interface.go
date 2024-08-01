package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type CategoryDao interface {
	UpdateCategoryType(ctx context.Context, category *model.Category) error
	GetCategoryById(ctx context.Context, id int) (*model.Category, error)
}
