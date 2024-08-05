package category

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

// ChangeCategoryStatus 修改分类状态
// @Param ctx context.Context 上下文
// @Param id int64 分类ID
// @Param status int 分类状态
// @Return err error 错误信息
func (dao *CategoryDaoImpl) ChangeCategoryStatus(ctx context.Context, id int64, status int) error {
	query := dao.DB.Model(&model.Category{})
	err := query.WithContext(ctx).Where("id = ?", id).Update("status = ?", status).Error
	return err
}

// CreateCategory 创建分类
// @Param ctx context.Context 上下文
// @Param category *model.Category 分类信息
// @Return err error 错误信息
func (dao *CategoryDaoImpl) CreateCategory(ctx context.Context, category *model.Category) error {
	err := dao.DB.WithContext(ctx).Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory 删除分类
// @Param ctx context.Context 上下文
// @Param id int64 分类ID
// @Return err error 错误信息
func (dao *CategoryDaoImpl) DeleteCategory(ctx context.Context, id int64) error {
	err := dao.DB.WithContext(ctx).Delete(&model.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCategoryByType 根据类型获取分类
// @Param ctx context.Context 上下文
// @Param categoryType int 分类类型
// @Return []model.Category 分类列表
// @Return error 错误信息
func (dao *CategoryDaoImpl) GetCategoryByType(ctx context.Context, categoryType int) ([]model.Category, error) {
	var categories []model.Category
	err := dao.DB.WithContext(ctx).Where("type = ?", categoryType).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
