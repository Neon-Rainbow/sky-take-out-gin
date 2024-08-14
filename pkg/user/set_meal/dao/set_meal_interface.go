package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type SetMealDaoInterface interface {
	// GetSetMealList 获取套餐列表
	GetSetMealList(ctx context.Context, categoryID int) (setMealList []model.Setmeal, err error)

	// GetSetMealDetail 获取套餐详情
	GetSetMealDetail(ctx context.Context, setMealID int) (setMealDetail []model.SetmealDish, err error)
}
