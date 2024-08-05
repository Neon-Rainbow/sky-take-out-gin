package service

import (
	"context"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/shop/DTO"
)

type ShopServiceInterface interface {
	GetShopStatus(ctx context.Context, req *DTO.GetShopStatusRequest) (*DTO.GetShopStatusResponse, *error2.ApiError)
	SetShopStatus(ctx context.Context, req *DTO.SetShopStatusRequest) (*DTO.SetShopStatusResponse, *error2.ApiError)
}
