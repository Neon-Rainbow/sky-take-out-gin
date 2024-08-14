package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	apiErrorModel "sky-take-out-gin/pkg/common/error"
)

type SetMealServiceInterface interface {
	// GetSetMealList 获取套餐列表
	GetSetMealList(ctx context.Context, categoryID int) (setMealList []model.Setmeal, apiError *apiErrorModel.ApiError)

	// GetSetMealDetail 获取套餐详情
	GetSetMealDetail(ctx context.Context, setMealID int) (setMealDetail []model.SetmealDish, apiError *apiErrorModel.ApiError)
}
