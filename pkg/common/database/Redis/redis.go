package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sky-take-out-gin/pkg/common/config"
)

var rdb *redis.Client

// InitRedis 初始化Redis
func InitRedis() error {
	cfg := config.GetConfig().RedisConfig
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试Redis是否连接成功
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetRedis 获取Redis
func GetRedis() *redis.Client {
	if rdb == nil {
		panic("Redis还未初始化")
	}
	return rdb
}
