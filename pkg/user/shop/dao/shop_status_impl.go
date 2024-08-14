package dao

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"sky-take-out-gin/pkg/common/database"
)

type ShopDaoImpl struct {
	db database.DatabaseInterface
}

// GetShopStatus 获取店铺状态
func (dao ShopDaoImpl) GetShopStatus(ctx context.Context) (int, error) {
	status, err := dao.db.GetRedis().Get(ctx, "shop_status").Int()
	if errors.Is(err, redis.Nil) {
		return -1, errors.New("在 Redis 中未找到 shop_status 字段, 请检查是否已经设置")
	}
	if err != nil {
		return -1, err
	}
	return status, nil
}

// NewShopDaoImpl 创建 ShopDaoImpl 实例
func NewShopDaoImpl(db database.DatabaseInterface) ShopDaoImpl {
	return ShopDaoImpl{db}
}
