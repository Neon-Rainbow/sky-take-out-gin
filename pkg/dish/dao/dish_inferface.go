package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type DishDaoInterface interface {
	CreateDish(ctx context.Context, dish model.Dish) error
	UpdateDish(ctx context.Context, dish model.Dish) error
	DeleteDish(ctx context.Context, ids []int64) error
	SearchDishByID(ctx context.Context, id int64) (*model.Dish, error)
	SearchDishByCategory(ctx context.Context, categoryID int64) ([]model.Dish, error)
	SearchDishByPage(ctx context.Context, categoryID int64, name string, status, page, pageSize int) (int, []model.Dish, error)
	ChangeDishStatus(ctx context.Context, id int64, status int) error
}
