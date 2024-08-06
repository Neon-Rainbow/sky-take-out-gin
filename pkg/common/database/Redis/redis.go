package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"sky-take-out-gin/pkg/common/config"
)

type RedisDB struct {
	redisClient *redis.Client
}

func (r RedisDB) InitRedis() error {
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
	r.redisClient = rdb
	return nil
}

func (r RedisDB) GetRedis() *redis.Client {
	if r.redisClient == nil {
		zap.L().Fatal("Redis还未初始化")
	}
	return r.redisClient
}

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
