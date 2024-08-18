package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"sky-take-out-gin/model/sql"
	paramModel "sky-take-out-gin/pkg/admin/category/DTO"
	controllerModel "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
)

// UpdateCategory 更新分类
func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, category *paramModel.AdminUpdateCategoryRequest) (*paramModel.AdminUpdateCategoryResponse, *controllerModel.ApiError) {
	existingCategory, err := service.CategoryDaoInfertace.GetCategoryById(ctx, category.ID)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryNotExist,
			Msg:  fmt.Sprintf("分类不存在, err: %v", err),
		}
	}

	err = service.cache.Invalidate(ctx, fmt.Sprintf("category_list_type:%v", existingCategory.Type))
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CacheInvalidateFailed,
			Msg:  fmt.Sprintf("缓存失效失败, err: %v", err),
		}
	}

	err = copier.CopyWithOption(existingCategory, category, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryUpdateFailed,
			Msg:  fmt.Sprintf("更新分类失败, err: %v", err),
		}
	}

	//err = convert.UpdateStructFields(category, existingCategory)
	//if err != nil {
	//	return nil, &controllerModel.ApiError{
	//		Code: code.CategoryUpdateFailed,
	//		Msg:  fmt.Sprintf("更新分类失败, err: %v", err),
	//	}
	//}

	err = service.UpdateCategoryType(ctx, existingCategory)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryUpdateFailed,
			Msg:  fmt.Sprintf("更新分类失败, err: %v", err),
		}
	}
	return nil, nil
}

// GetCategoryPage 分页查询分类
// @Param ctx context.Context 上下文
func (service *CategoryServiceImpl) GetCategoryPage(ctx context.Context, req *paramModel.AdminCategoryPageRequest) (*paramModel.AdminCategoryPageResponse, *controllerModel.ApiError) {
	// 分页查询分类
	categories, total, err := service.CategoryDaoInfertace.GetCategoryPage(ctx, req.Name, req.Page, req.PageSize, req.Type)
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

// ChangeCategoryStatus 启用、禁用分类
func (service *CategoryServiceImpl) ChangeCategoryStatus(ctx context.Context, p *paramModel.AdminChangeCategoryStatusRequest) (*paramModel.AdminChangeCategoryStatusResponse, *controllerModel.ApiError) {
	err := service.CategoryDaoInfertace.ChangeCategoryStatus(ctx, p.ID, p.Status)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryUpdateFailed,
			Msg:  fmt.Sprintf("更新分类失败, err: %v", err),
		}
	}

	err = service.cache.InvalidatePattern(ctx, fmt.Sprintf("category_list_type:%v", "*"))
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CacheInvalidateFailed,
			Msg:  fmt.Sprintf("缓存失效失败, err: %v", err),
		}
	}

	return nil, nil
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, p *paramModel.AdminCreateCategoryRequest) (*paramModel.AdminCreateCategoryResponse, *controllerModel.ApiError) {
	userID, ok := ctx.Value("userID").(uint)
	if !ok {
		return nil, &controllerModel.ApiError{
			Code: code.ParamError,
			Msg:  "无效的用户ID",
		}
	}
	category := &model.Category{
		Type:       p.Type,
		Name:       p.Name,
		Sort:       p.Sort,
		Status:     1,
		CreateUser: userID,
		UpdateUser: userID,
	}
	err := service.CategoryDaoInfertace.CreateCategory(ctx, category)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryCreateFailed,
			Msg:  fmt.Sprintf("新增分类失败, err: %v", err),
		}
	}
	return nil, nil
}

// DeleteCategory 删除分类
// @Param ctx context.Context 上下文
// @Param p *paramModel.AdminDeleteCategoryRequest 删除分类请求
// @Return *paramModel.AdminDeleteCategoryResponse 删除分类响应
// @Return *controllerModel.ApiError 错误信息
func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, p *paramModel.AdminDeleteCategoryRequest) (*paramModel.AdminDeleteCategoryResponse, *controllerModel.ApiError) {
	err := service.CategoryDaoInfertace.DeleteCategory(ctx, p.ID)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryDeleteFailed,
			Msg:  fmt.Sprintf("删除分类失败, err: %v", err),
		}
	}
	err = service.cache.InvalidatePattern(ctx, fmt.Sprintf("category_list_type:%v", "*"))
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CacheInvalidateFailed,
			Msg:  fmt.Sprintf("缓存失效失败, err: %v", err),
		}
	}

	return nil, nil
}

// GetCategoryByType 根据类型获取分类
// @Param ctx context.Context 上下文
// @Param p *paramModel.AdminGetCategoryListByTypeRequest 根据类型查询分类请求

func (service *CategoryServiceImpl) GetCategoryByType(ctx context.Context, p *paramModel.AdminGetCategoryListByTypeRequest) (*paramModel.AdminGetCategoryListByTypeResponse, *controllerModel.ApiError) {
	categories, err := service.CategoryDaoInfertace.GetCategoryByType(ctx, p.Type)
	if err != nil {
		return nil, &controllerModel.ApiError{
			Code: code.CategoryGetFailed,
			Msg:  fmt.Sprintf("获取分类失败, err: %v", err),
		}
	}
	return &paramModel.AdminGetCategoryListByTypeResponse{
		CategoryList: categories,
	}, nil
}
