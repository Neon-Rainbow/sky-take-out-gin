package category

import (
	"context"
	model2 "sky-take-out-gin/model"
	paramModel "sky-take-out-gin/model/param/admin/category"
	model "sky-take-out-gin/model/sql"
)

type CategoryService interface {
	UpdateCategory(ctx context.Context, category *model.Category) *model2.ApiError
	GetCategoryPage(ctx context.Context, req *paramModel.AdminCategoryPageRequest) (*paramModel.AdminCategoryPageResponse, *model2.ApiError)
	ChangeCategoryStatus(ctx context.Context, p *paramModel.AdminChangeCategoryStatusRequest) *model2.ApiError
	CreateCategory(ctx context.Context, p *paramModel.AdminCreateCategoryRequest) *model2.ApiError
}
