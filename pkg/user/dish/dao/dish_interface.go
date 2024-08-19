package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type DishDaoInterface interface {
	GetDishByCategoryId(ctx context.Context, categoryID uint) (dishes []model.Dish, err error)
	GetDishByDishID(ctx context.Context, dishID uint) (dish model.Dish, err error)
}
