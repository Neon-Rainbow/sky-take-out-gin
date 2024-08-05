package dish

import (
	"context"
	"sky-take-out-gin/model"
	paramModel "sky-take-out-gin/model/param/admin/dish"
)

type DishServiceInterface interface {
	UpdateDish(ctx context.Context, req *paramModel.UpdateDishRequest) (resp *paramModel.UpdateDishResponse, apiError *model.ApiError)
	DeleteDish(ctx context.Context, req *paramModel.DeleteDishRequest) (resp *paramModel.DeleteDishResponse, apiError *model.ApiError)
	AddDish(ctx context.Context, req *paramModel.AddDishRequest) (resp *paramModel.AddDishResponse, apiError *model.ApiError)
	SearchDishByID(ctx context.Context, req *paramModel.SearchDishByIDRequest) (resp *paramModel.SearchDishByIDResponse, apiError *model.ApiError)
	SearchDishByCategory(ctx context.Context, req *paramModel.SearchDishByCategoryRequest) (resp *paramModel.SearchDishByCategoryResponse, apiError *model.ApiError)
	SearchDishByPage(ctx context.Context, req *paramModel.SearchDishByPageRequest) (resp *paramModel.SearchDishByPageResponse, apiError *model.ApiError)
	ChangeDishStatus(ctx context.Context, req *paramModel.ChangeDishStatusRequest) (resp *paramModel.ChangeDishStatusResponse, apiError *model.ApiError)
}
