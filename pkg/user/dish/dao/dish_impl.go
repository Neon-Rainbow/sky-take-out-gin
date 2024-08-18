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

func (d *DishDaoImpl) GetDishById(ctx context.Context, categoryID uint) (dishes []model.Dish, err error) {
	err = d.db.GetDB().
		WithContext(ctx).
		Preload("DishFlavors"). // 预加载关联表
		Where("category_id = ?", categoryID).
		Find(&dishes).
		Error
	return
}
