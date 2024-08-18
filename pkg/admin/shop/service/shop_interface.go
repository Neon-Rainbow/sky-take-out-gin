package service

import (
	"context"
	"sky-take-out-gin/pkg/admin/shop/DTO"
	error2 "sky-take-out-gin/pkg/common/api_error"
)

type ShopServiceInterface interface {
	GetShopStatus(ctx context.Context, req *DTO.GetShopStatusRequest) (*DTO.GetShopStatusResponse, *error2.ApiError)
	SetShopStatus(ctx context.Context, req *DTO.SetShopStatusRequest) (*DTO.SetShopStatusResponse, *error2.ApiError)
}
