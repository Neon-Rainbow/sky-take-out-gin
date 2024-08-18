package service

import (
	"context"
	apiErrorModel "sky-take-out-gin/pkg/common/api_error"
)

type ShopStatusServiceInterface interface {
	// GetShopStatus 获取店铺状态
	GetShopStatus(ctx context.Context) (status int, apiError *apiErrorModel.ApiError)
}
