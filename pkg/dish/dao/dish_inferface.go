package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type DishDaoInterface interface {
	CreateDish(ctx context.Context, dish model.Dish) error
	UpdateDish(ctx context.Context, dish model.Dish) error
	DeleteDish(ctx context.Context, ids []uint) error
	SearchDishByID(ctx context.Context, id uint) (*model.Dish, error)
	SearchDishByCategory(ctx context.Context, categoryID uint) ([]model.Dish, error)
	SearchDishByPage(ctx context.Context, categoryID uint, name string, status, page, pageSize int) (total int, records []model.Dish, err error)
	ChangeDishStatus(ctx context.Context, id uint, status int) error
}
