package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type SetMealDaoImpl struct {
	db database.DatabaseInterface
}

// GetSetMealList 获取套餐列表
func (s SetMealDaoImpl) GetSetMealList(ctx context.Context, categoryID int) (setMealList []model.Setmeal, err error) {
	var setMeals []model.Setmeal
	err = s.db.GetDB().WithContext(ctx).Where("category_id = ?", categoryID).Find(&setMeals).Error
	if err != nil {
		return nil, err
	}
	return setMeals, nil
}

// GetSetMealDetail 获取套餐详情
func (s SetMealDaoImpl) GetSetMealDetail(ctx context.Context, setMealID int) (setMealDetail []model.SetmealDish, err error) {
	var setMealDishes []model.SetmealDish
	err = s.db.GetDB().WithContext(ctx).Where("setmeal_id = ?", setMealID).Find(&setMealDishes).Error
	if err != nil {
		return nil, err
	}
	return setMealDishes, nil
}

// NewSetMealDaoImpl 实例化SetMealDaoImpl
func NewSetMealDaoImpl(db database.DatabaseInterface) *SetMealDaoImpl {
	return &SetMealDaoImpl{db: db}
}
