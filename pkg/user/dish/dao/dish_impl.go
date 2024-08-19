package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type DishDaoImpl struct {
	db database.DatabaseInterface
}

func NewDishDao(db database.DatabaseInterface) *DishDaoImpl {
	return &DishDaoImpl{db: db}
}

func (d *DishDaoImpl) GetDishByCategoryId(ctx context.Context, categoryID uint) (dishes []model.Dish, err error) {
	err = d.db.GetDB().
		WithContext(ctx).
		Preload("DishFlavors"). // 预加载关联表
		Where("category_id = ?", categoryID).
		Find(&dishes).
		Error
	return
}

func (d *DishDaoImpl) GetDishByDishID(ctx context.Context, dishID uint) (dish model.Dish, err error) {
	err = d.db.GetDB().
		WithContext(ctx).
		Preload("DishFlavors"). // 预加载关联表
		Where("id = ?", dishID).
		First(&dish).
		Error
	return
}
