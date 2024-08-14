package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/user/set_meal/dao"
)

type SetMealServiceImpl struct {
	dao dao.SetMealDaoInterface
}

// GetSetMealList 获取套餐列表
func (s SetMealServiceImpl) GetSetMealList(ctx context.Context, categoryID int) (setMealList []model.Setmeal, apiError *error2.ApiError) {
	var err error
	setMealList, err = s.dao.GetSetMealList(ctx, categoryID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SetMealGetListError,
			Msg:  err.Error(),
		}
	}
	return setMealList, nil
}

// GetSetMealDetail 获取套餐详情
func (s SetMealServiceImpl) GetSetMealDetail(ctx context.Context, setMealID int) (setMealDetail []model.SetmealDish, apiError *error2.ApiError) {
	var err error
	setMealDetail, err = s.dao.GetSetMealDetail(ctx, setMealID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SetMealGetDetailError,
			Msg:  err.Error(),
		}
	}
	return setMealDetail, nil
}

// NewSetMealServiceImpl 实例化SetMealServiceImpl
func NewSetMealServiceImpl(dao dao.SetMealDaoInterface) *SetMealServiceImpl {
	return &SetMealServiceImpl{dao: dao}
}
