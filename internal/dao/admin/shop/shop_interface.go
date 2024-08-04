package shop

import "context"

type ShopDaoInterface interface {
	GetShopStatus(ctx context.Context) (int, error)
	SetShopStatus(ctx context.Context, status int) error
}
