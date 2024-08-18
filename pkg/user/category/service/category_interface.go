package service

import (
	"context"
	apiErrorModel "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/user/category/DTO"
)

type CategoryServiceInterface interface {
	// GetCategoryList 获取分类列表
	GetCategoryList(ctx context.Context, req *DTO.CategoryRequestDTO) (*DTO.CategoryResponseDTO, *apiErrorModel.ApiError)
}
