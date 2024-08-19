package service

import (
	"context"
	"fmt"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/user/dish/dao"
	"time"
)

type DishServiceImpl struct {
	dao   dao.DishDaoInterface
	cache cache.RedisCacheInterface
}

func (service DishServiceImpl) GetDishByID(ctx context.Context, categoryID uint) (dishes []model.Dish, apiError *api_error.ApiError) {
	var err error

	cacheID := fmt.Sprintf("dish_by_category_id:%d", categoryID)

	// 从缓存中获取数据
	err = service.cache.GetOrSet(
		ctx,
		cacheID,
		time.Hour,
		&dishes,
		func(ctx context.Context, args ...interface{}) (interface{}, error) {
			return service.dao.GetDishByCategoryId(ctx, categoryID)
		},
	)
	if err != nil {
		apiError = api_error.NewApiError(code.GetDishByIdError, err)
		return nil, apiError
	}

	return dishes, nil
}

func NewDishService(dao dao.DishDaoInterface, cache cache.RedisCacheInterface) DishServiceImpl {
	return DishServiceImpl{dao, cache}
}
