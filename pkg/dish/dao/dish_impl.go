package dao

import (
	"context"
	"gorm.io/gorm"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type DishDaoImpl struct {
	db database.DatabaseInterface
}

func (dao DishDaoImpl) CreateDish(ctx context.Context, dish model.Dish) error {
	err := dao.db.GetDB().WithContext(ctx).Create(&dish).Error
	return err
}

func (dao DishDaoImpl) CreateDishWithTransaction(ctx context.Context, tx *gorm.DB, dish *model.Dish) error {
	err := tx.WithContext(ctx).Create(&dish).Error
	return err
}

func (dao DishDaoImpl) UpdateDish(ctx context.Context, dish model.Dish) error {
	err := dao.db.GetDB().WithContext(ctx).Save(&dish).Error
	return err
}

func (dao DishDaoImpl) UpdateDishWithTransaction(ctx context.Context, tx *gorm.DB, dish *model.Dish) error {
	err := tx.WithContext(ctx).Save(&dish).Error
	return err
}

func (dao DishDaoImpl) DeleteDish(ctx context.Context, ids []uint) error {
	err := dao.db.GetDB().WithContext(ctx).Where("id in ?", ids).Delete(&model.Dish{}).Error
	return err
}

func (dao DishDaoImpl) SearchDishByID(ctx context.Context, id uint) (*model.Dish, error) {
	var dish *model.Dish
	err := dao.db.GetDB().WithContext(ctx).Preload("DishFlavors").Where("id = ?", id).First(&dish).Error
	if err != nil {
		return nil, err
	}
	return dish, err
}

func (dao DishDaoImpl) SearchDishByCategory(ctx context.Context, categoryID uint) ([]model.Dish, error) {
	var dishes []model.Dish
	err := dao.db.GetDB().WithContext(ctx).Where("category_id = ?", categoryID).Find(&dishes).Error
	return dishes, err
}

func (dao DishDaoImpl) SearchDishByPage(ctx context.Context, categoryID uint, name string, status, page, pageSize int) (int, []model.Dish, error) {
	var dishes []model.Dish
	var count int64
	db := dao.db.GetDB().Model(&model.Dish{}).WithContext(ctx)
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

func (dao DishDaoImpl) ChangeDishStatus(ctx context.Context, id uint, status int) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Dish{}).Where("id = ?", id).Update("status", status).Error
	return err
}

func (dao DishDaoImpl) UpdateDishFlavor(ctx context.Context, flavor model.DishFlavor) error {
	err := dao.db.GetDB().WithContext(ctx).Save(&flavor).Error
	return err
}

func (dao DishDaoImpl) CreateDishFlavor(ctx context.Context, flavor model.DishFlavor) error {
	err := dao.db.GetDB().WithContext(ctx).Create(&flavor).Error
	return err
}

func (dao DishDaoImpl) CreateDishFlavorWithTransaction(ctx context.Context, tx *gorm.DB, flavor model.DishFlavor) error {
	err := tx.WithContext(ctx).Create(&flavor).Error
	return err
}

func (dao DishDaoImpl) DeleteDishFlavorsByDishIDWithTransaction(ctx context.Context, tx *gorm.DB, dishID uint) error {
	err := tx.WithContext(ctx).Where("dish_id = ?", dishID).Delete(&model.DishFlavor{}).Error
	return err
}

func (dao DishDaoImpl) BeginTransaction() *gorm.DB {
	return dao.db.GetDB().Begin()
}

func NewDishDaoImpl(db database.DatabaseInterface) *DishDaoImpl {
	return &DishDaoImpl{db}
}
