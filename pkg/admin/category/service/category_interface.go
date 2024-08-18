package service

import (
	"context"
	paramModel "sky-take-out-gin/pkg/admin/category/DTO"
	model2 "sky-take-out-gin/pkg/common/api_error"
)

type CategoryService interface {
	UpdateCategory(ctx context.Context, category *paramModel.AdminUpdateCategoryRequest) (*paramModel.AdminUpdateCategoryResponse, *model2.ApiError)
	GetCategoryPage(ctx context.Context, req *paramModel.AdminCategoryPageRequest) (*paramModel.AdminCategoryPageResponse, *model2.ApiError)
	ChangeCategoryStatus(ctx context.Context, p *paramModel.AdminChangeCategoryStatusRequest) (*paramModel.AdminChangeCategoryStatusResponse, *model2.ApiError)
	CreateCategory(ctx context.Context, p *paramModel.AdminCreateCategoryRequest) (*paramModel.AdminCreateCategoryResponse, *model2.ApiError)
	DeleteCategory(ctx context.Context, p *paramModel.AdminDeleteCategoryRequest) (*paramModel.AdminDeleteCategoryResponse, *model2.ApiError)
	GetCategoryByType(ctx context.Context, p *paramModel.AdminGetCategoryListByTypeRequest) (*paramModel.AdminGetCategoryListByTypeResponse, *model2.ApiError)
}
