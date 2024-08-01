package category

import (
	"context"
	model2 "sky-take-out-gin/model"
	model "sky-take-out-gin/model/sql"
)

type CategoryService interface {
	UpdateCategory(ctx context.Context, category *model.Category) *model2.ApiError
}
