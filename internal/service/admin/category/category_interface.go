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
	ChangeCategoryStatus(ctx context.Context, p *paramModel.AdminChangeCategoryStatusRequest) (*paramModel.AdminChangeCategoryStatusResponse, *model2.ApiError)
	CreateCategory(ctx context.Context, p *paramModel.AdminCreateCategoryRequest) (*paramModel.AdminCreateCategoryResponse, *model2.ApiError)
	DeleteCategory(ctx context.Context, p *paramModel.AdminDeleteCategoryRequest) (*paramModel.AdminDeleteCategoryResponse, *model2.ApiError)
	GetCategoryByType(ctx context.Context, p *paramModel.AdminGetCategoryListByTypeRequest) (*paramModel.AdminGetCategoryListByTypeResponse, *model2.ApiError)
}
