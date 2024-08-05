package shop

import (
	"context"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/shop"
)

type ShopServiceInterface interface {
	GetShopStatus(ctx context.Context, req *shop.GetShopStatusRequest) (*shop.GetShopStatusResponse, *model.ApiError)
	SetShopStatus(ctx context.Context, req *shop.SetShopStatusRequest) (*shop.SetShopStatusResponse, *model.ApiError)
}
