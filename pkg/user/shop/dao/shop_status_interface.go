package dao

import "context"

type ShopStatusDaoInterface interface {
	// GetShopStatus 获取店铺状态
	GetShopStatus(ctx context.Context) (int, error)
}
