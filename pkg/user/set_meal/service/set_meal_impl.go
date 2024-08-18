package service

import (
	"context"
	"fmt"
	model "sky-take-out-gin/model/sql"
	error2 "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/user/set_meal/dao"
	"time"
)

type SetMealServiceImpl struct {
	dao   dao.SetMealDaoInterface
	cache cache.RedisCacheInterface
}

// GetSetMealList 获取套餐列表
func (s SetMealServiceImpl) GetSetMealList(ctx context.Context, categoryID int) (setMealList []model.Setmeal, apiError *error2.ApiError) {
	cacheKey := fmt.Sprintf("set_meal_list_category_id:%v", categoryID)
	err := s.cache.GetOrSet(
		ctx,
		cacheKey,
		time.Hour,
		&setMealList,
		func(ctx context.Context, args ...interface{}) (interface{}, error) {
			return s.dao.GetSetMealList(ctx, categoryID)
		},
	)

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
	cacheKey := fmt.Sprintf("set_meal_detail_id:%v", setMealID)
	err := s.cache.GetOrSet(
		ctx,
		cacheKey,
		time.Hour,
		&setMealDetail,
		func(ctx context.Context, args ...interface{}) (interface{}, error) {
			return s.dao.GetSetMealDetail(ctx, setMealID)
		},
	)

	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SetMealGetDetailError,
			Msg:  err.Error(),
		}
	}
	return setMealDetail, nil
}

// NewSetMealServiceImpl 实例化SetMealServiceImpl
func NewSetMealServiceImpl(dao dao.SetMealDaoInterface, cache cache.RedisCacheInterface) *SetMealServiceImpl {
	return &SetMealServiceImpl{dao: dao, cache: cache}
}
