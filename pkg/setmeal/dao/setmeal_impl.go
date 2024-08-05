package dao

import (
	"context"
	"gorm.io/gorm"
	sqlModel "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database/MySQL"
)

type SetmealDAOImpl struct {
	db *gorm.DB
}

// CreateSetmeal 创建套餐
// @Param ctx context.Context 上下文
// @Param setmeal *model.Setmeal 套餐信息
// @Return error 错误信息
func (dao SetmealDAOImpl) CreateSetmeal(ctx context.Context, setmeal *sqlModel.Setmeal) error {
	return dao.db.WithContext(ctx).Create(setmeal).Error
}

// GetSetmealByID 根据ID获取套餐
// @Param ctx context.Context 上下文
// @Param id int64 套餐ID
// @Return *model.Setmeal 套餐信息
func (dao SetmealDAOImpl) GetSetmealByID(ctx context.Context, id int64) (*sqlModel.Setmeal, error) {
	var setmeal sqlModel.Setmeal
	err := dao.db.WithContext(ctx).Preload("SetmealDishes").First(&setmeal, id).Error
	return &setmeal, err
}

// GetSetmealPage 获取套餐分页
// @Param ctx context.Context 上下文
// @Param page int 页码
// @Param pageSize int 每页数量
// @Return []sqlModel.Setmeal 套餐列表
// @Return error 错误信息
func (dao SetmealDAOImpl) GetSetmealPage(ctx context.Context, CategoryID int64, page, pageSize int) ([]sqlModel.Setmeal, error) {
	var setmeals []sqlModel.Setmeal
	offset := (page - 1) * pageSize
	if err := dao.db.WithContext(ctx).Offset(offset).Limit(pageSize).Where("category_id = ?", CategoryID).Find(&setmeals).Error; err != nil {
		return nil, err
	}
	return setmeals, nil
}

func (dao SetmealDAOImpl) SearchSetmeals(ctx context.Context, condition string, args ...interface{}) ([]sqlModel.Setmeal, error) {
	var setmeals []sqlModel.Setmeal
	if err := dao.db.WithContext(ctx).Preload("SetmealDishes").Where(condition, args...).Find(&setmeals).Error; err != nil {
		return nil, err
	}
	return setmeals, nil
}

// UpdateSetmeal 更新套餐
func (dao SetmealDAOImpl) UpdateSetmeal(ctx context.Context, setmeal *sqlModel.Setmeal) error {
	return dao.db.WithContext(ctx).Save(setmeal).Error
}

// UpdateSetmealStatus 更新套餐状态
func (dao SetmealDAOImpl) UpdateSetmealStatus(ctx context.Context, id int64, status int) error {
	return dao.db.WithContext(ctx).Model(&sqlModel.Setmeal{}).Where("id = ?", id).Update("status", status).Error
}

// DeleteSetmeals 删除套餐
func (dao SetmealDAOImpl) DeleteSetmeals(ctx context.Context, ids []int64) error {
	return dao.db.WithContext(ctx).Preload("SetmealDishes").Where("id in (?)", ids).Delete(&sqlModel.Setmeal{}).Error
}

// GetSetmeals 获取套餐列表
func (dao SetmealDAOImpl) GetSetmeals(ctx context.Context, categoryID int64, name string, status int, offset int, limit int) ([]sqlModel.Setmeal, int64, error) {
	var setmeals []sqlModel.Setmeal
	var total int64
	db := dao.db.WithContext(ctx).Preload("SetmealDishes")
	if categoryID != 0 {
		db = db.Where("category_id = ?", categoryID)
	}
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if status != -1 {
		db = db.Where("status = ?", status)
	}
	if err := db.Offset(offset).Limit(limit).Find(&setmeals).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return setmeals, total, nil
}

// SetmealRawSQL 原生SQL查询
func (dao SetmealDAOImpl) SetmealRawSQL(ctx context.Context, sql string, values ...interface{}) ([]sqlModel.Setmeal, error) {
	var setmeals []sqlModel.Setmeal
	err := dao.db.WithContext(ctx).Raw(sql, values...).Scan(&setmeals).Error
	if err != nil {
		return nil, err
	}
	return setmeals, nil
}

// NewSetmealDAOImpl 实例化SetmealDAOImpl
func NewSetmealDAOImpl() *SetmealDAOImpl {
	return &SetmealDAOImpl{MySQL.GetDB()}
}
