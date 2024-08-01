package category

import (
	"context"
	"fmt"
	"sky-take-out-gin/code"
	controllerModel "sky-take-out-gin/model"
	paramModel "sky-take-out-gin/model/param/admin/category"
	"sky-take-out-gin/model/sql"
)

// UpdateCategory 更新分类
func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, category *model.Category) *controllerModel.ApiError {
	existingCategory, err := service.CategoryDaoImpl.GetCategoryById(ctx, category.ID)
	if err != nil {
		return &controllerModel.ApiError{
			Code: code.CategoryNotExist,
			Msg:  fmt.Sprintf("分类不存在, err: %v", err),
		}
	}

	existingCategory.Name = category.Name
	existingCategory.Type = category.Type
	existingCategory.Sort = category.Sort
	existingCategory.Status = category.Status
	existingCategory.UpdateTime = category.UpdateTime
	existingCategory.UpdateUser = category.UpdateUser

	err = service.CategoryDaoImpl.UpdateCategoryType(ctx, existingCategory)
	if err != nil {
		return &controllerModel.ApiError{
			Code: code.CategoryUpdateFailed,
			Msg:  fmt.Sprintf("更新分类失败, err: %v", err),
		}
	}
	return nil
}

// GetCategoryPage 分页查询分类
// @Param ctx context.Context 上下文
func (service *CategoryServiceImpl) GetCategoryPage(ctx context.Context, req *paramModel.AdminCategoryPageRequest) (*paramModel.AdminCategoryPageResponse, *controllerModel.ApiError) {
	// 分页查询分类
	categories, total, err := service.CategoryDaoImpl.GetCategoryPage(ctx, req.Name, req.Page, req.PageSize, req.Type)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryGetFailed,
			Msg:  fmt.Sprintf("获取分类失败, err: %v", err),
		}
	}

	// 将分类列表转换为响应结构
	var records []paramModel.CategoryRecord
	for _, category := range categories {
		records = append(records, paramModel.CategoryRecord{
			ID:     category.ID,
			Type:   category.Type,
			Name:   category.Name,
			Sort:   category.Sort,
			Status: category.Status,
		})
	}

	return &paramModel.AdminCategoryPageResponse{
		Records: records,
		Total:   total,
	}, nil
}
