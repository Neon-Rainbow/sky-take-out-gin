package shop

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"sky-take-out-gin/utils/Redis"
	"time"
)

type ShopDaoImpl struct {
	redisClient *redis.Client
}

// GetShopStatus 获取店铺状态
func (dao ShopDaoImpl) GetShopStatus(ctx context.Context) (int, error) {
	status, err := dao.redisClient.Get(ctx, "shop_status").Int()
	if errors.Is(err, redis.Nil) {
		return -1, errors.New("在 Redis 中未找到 shop_status 字段, 请检查是否已经设置")
	}
	if err != nil {
		return -1, err
	}
	return status, nil
}

// SetShopStatus 设置店铺状态
func (dao ShopDaoImpl) SetShopStatus(ctx context.Context, status int) error {
	err := dao.redisClient.Set(ctx, "shop_status", status, time.Hour*24).Err()
	return err
}

func NewShopDaoImpl() ShopDaoImpl {
	return ShopDaoImpl{redisClient: Redis.GetRedis()}
}
