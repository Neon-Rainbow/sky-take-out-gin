package service

import (
	"context"
	"fmt"
	model "sky-take-out-gin/model/sql"
	error2 "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/user/category/DTO"
	"sky-take-out-gin/pkg/user/category/dao"
	"time"
)

type CategoryServiceImpl struct {
	dao.CategoryDaoInterface
	cache cache.RedisCacheInterface // 新增缓存接口字段
}

// GetCategoryList 获取分类列表
func (service CategoryServiceImpl) GetCategoryList(ctx context.Context, req *DTO.CategoryRequestDTO) (*DTO.CategoryResponseDTO, *error2.ApiError) {
	// 定义缓存键
	cacheKey := fmt.Sprintf("category_list_type:%v", req.Type)

	// 尝试从缓存中获取数据
	var categories []model.Category
	err := service.cache.GetOrSet(
		ctx,
		cacheKey,
		time.Hour,
		&categories,
		func(ctx context.Context, args ...interface{}) (interface{}, error) {
			// 如果缓存中没有数据，从数据库获取
			return service.CategoryDaoInterface.GetCategoryList(ctx, req.Type)
		},
	)

	if err != nil {
		return nil, &error2.ApiError{
			Code: code.CategoryGetListError,
			Msg:  err.Error(),
		}
	}

	// 构建响应对象
	resp := &DTO.CategoryResponseDTO{
		Categories: categories,
	}

	return resp, nil
}

// NewCategoryService 创建分类服务
func NewCategoryService(dao dao.CategoryDaoInterface, cache cache.RedisCacheInterface) *CategoryServiceImpl {
	return &CategoryServiceImpl{dao, cache}
}
