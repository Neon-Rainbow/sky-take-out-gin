package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type DishDaoInterface interface {
	GetDishById(ctx context.Context, categoryID uint) (dishes []model.Dish, err error)
}
