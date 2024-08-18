package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	error2 "sky-take-out-gin/pkg/common/api_error"
)

type DishServiceInterface interface {
	GetDishByID(ctx context.Context, categoryID uint) (dishes []model.Dish, apiError *error2.ApiError)
}
