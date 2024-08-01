package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

// UpdateCategoryType 更新分类类型
func (dao *CategoryDaoImpl) UpdateCategoryType(ctx context.Context, category *model.Category) error {
	err := dao.DB.WithContext(ctx).Save(category).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCategoryById 根据ID获取分类
func (dao *CategoryDaoImpl) GetCategoryById(ctx context.Context, id int64) (*model.Category, error) {
	category := &model.Category{}
	err := dao.DB.WithContext(ctx).First(category, id).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetCategoryPage 分页查询分类
// @Param ctx context.Context 上下文
// @Param name string 分类名称
// @Param page int 页码
// @Param pageSize int 每页记录数
// @Param categoryType int 分类类型
// @Return categories []model.Category 分类列表
// @Return total int64 总记录数
// @Return err error 错误信息
func (dao *CategoryDaoImpl) GetCategoryPage(ctx context.Context, name string, page, pageSize, categoryType int) (categories []model.Category, total int64, err error) {
	query := dao.DB.Model(&model.Category{})

	if name != "" {
		query = query.WithContext(ctx).Where("name LIKE ?", "%"+name+"%")
	}

	if categoryType != 0 {
		query = query.WithContext(ctx).Where("type = ?", categoryType)
	}

	if err = query.WithContext(ctx).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = query.WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}
