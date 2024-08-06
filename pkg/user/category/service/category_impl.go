package service

import (
	"context"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/user/category/DTO"
	"sky-take-out-gin/pkg/user/category/dao"
)

type CategoryServiceImpl struct {
	dao.CategoryDaoInterface
}

// GetCategoryList 获取分类列表
func (service CategoryServiceImpl) GetCategoryList(ctx context.Context, req *DTO.CategoryRequestDTO) (*DTO.CategoryResponseDTO, *error2.ApiError) {
	categories, err := service.CategoryDaoInterface.GetCategoryList(ctx, req.Type)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.CategoryGetListError,
			Msg:  err.Error(),
		}
	}
	resp := &DTO.CategoryResponseDTO{
		Categories: categories,
	}
	return resp, nil
}

func NewCategoryService(dao dao.CategoryDaoInterface) *CategoryServiceImpl {
	return &CategoryServiceImpl{dao}
}
