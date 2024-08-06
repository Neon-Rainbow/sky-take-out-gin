package dao

import (
	"context"
	"gorm.io/gorm"
	model "sky-take-out-gin/model/sql"
)

type DishDaoInterface interface {
	CreateDish(ctx context.Context, dish model.Dish) error
	CreateDishWithTransaction(ctx context.Context, tx *gorm.DB, dish *model.Dish) error

	UpdateDish(ctx context.Context, dish model.Dish) error
	UpdateDishWithTransaction(ctx context.Context, tx *gorm.DB, dish *model.Dish) error

	DeleteDish(ctx context.Context, ids []uint) error
	SearchDishByID(ctx context.Context, id uint) (*model.Dish, error)
	SearchDishByCategory(ctx context.Context, categoryID uint) ([]model.Dish, error)
	SearchDishByPage(ctx context.Context, categoryID uint, name string, status, page, pageSize int) (total int, records []model.Dish, err error)
	ChangeDishStatus(ctx context.Context, id uint, status int) error
	UpdateDishFlavor(ctx context.Context, flavor model.DishFlavor) error
	CreateDishFlavor(ctx context.Context, flavor model.DishFlavor) error
	CreateDishFlavorWithTransaction(ctx context.Context, tx *gorm.DB, flavor model.DishFlavor) error

	DeleteDishFlavorsByDishIDWithTransaction(ctx context.Context, tx *gorm.DB, dishID uint) error

	BeginTransaction() *gorm.DB
}
