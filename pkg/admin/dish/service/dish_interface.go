package service

import (
	"context"
	paramModel "sky-take-out-gin/pkg/admin/dish/DTO"
	error2 "sky-take-out-gin/pkg/common/api_error"
)

type DishServiceInterface interface {
	UpdateDish(ctx context.Context, req *paramModel.UpdateDishRequest) (resp *paramModel.UpdateDishResponse, apiError *error2.ApiError)
	DeleteDish(ctx context.Context, req *paramModel.DeleteDishRequest) (resp *paramModel.DeleteDishResponse, apiError *error2.ApiError)
	AddDish(ctx context.Context, req *paramModel.AddDishRequest) (resp *paramModel.AddDishResponse, apiError *error2.ApiError)
	SearchDishByID(ctx context.Context, req *paramModel.SearchDishByIDRequest) (resp *paramModel.SearchDishByIDResponse, apiError *error2.ApiError)
	SearchDishByCategory(ctx context.Context, req *paramModel.SearchDishByCategoryRequest) (resp *paramModel.SearchDishByCategoryResponse, apiError *error2.ApiError)
	SearchDishByPage(ctx context.Context, req *paramModel.SearchDishByPageRequest) (resp *paramModel.SearchDishByPageResponse, apiError *error2.ApiError)
	ChangeDishStatus(ctx context.Context, req *paramModel.ChangeDishStatusRequest) (resp *paramModel.ChangeDishStatusResponse, apiError *error2.ApiError)
}
