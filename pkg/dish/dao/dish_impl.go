package dao

import (
	"context"
	"gorm.io/gorm"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database/MySQL"
)

type DishDaoImpl struct {
	db *gorm.DB
}

func (dao DishDaoImpl) CreateDish(ctx context.Context, dish model.Dish) error {
	err := dao.db.WithContext(ctx).Create(&dish).Error
	return err
}

func (dao DishDaoImpl) UpdateDish(ctx context.Context, dish model.Dish) error {
	err := dao.db.WithContext(ctx).Save(&dish).Error
	return err
}

func (dao DishDaoImpl) DeleteDish(ctx context.Context, ids []int64) error {
	err := dao.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Dish{}).Error
	return err
}

func (dao DishDaoImpl) SearchDishByID(ctx context.Context, id int64) (*model.Dish, error) {
	var dish *model.Dish
	err := dao.db.WithContext(ctx).Preload("DishFlavors").Where("id = ?", id).First(&dish).Error
	if err != nil {
		return nil, err
	}
	return dish, err
}

func (dao DishDaoImpl) SearchDishByCategory(ctx context.Context, categoryID int64) ([]model.Dish, error) {
	var dishes []model.Dish
	err := dao.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&dishes).Error
	return dishes, err
}

func (dao DishDaoImpl) SearchDishByPage(ctx context.Context, categoryID int64, name string, status, page, pageSize int) (int, []model.Dish, error) {
	var dishes []model.Dish
	var count int64
	db := dao.db.Model(&model.Dish{}).WithContext(ctx)
	if categoryID != 0 {
		db = db.Where("category_id = ?", categoryID)
	}
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if status != -1 {
		db = db.Where("status = ?", status)
	}
	err := db.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dishes).Error
	return int(count), dishes, err
}

func (dao DishDaoImpl) ChangeDishStatus(ctx context.Context, id int64, status int) error {
	err := dao.db.WithContext(ctx).Model(&model.Dish{}).Where("id = ?", id).Update("status", status).Error
	return err
}

func NewDishDaoImpl() *DishDaoImpl {
	return &DishDaoImpl{MySQL.GetDB()}
}
