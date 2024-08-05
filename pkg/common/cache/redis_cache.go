package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"sky-take-out-gin/pkg/common/database/Redis"
	"time"
)

// Cache 结构体
type Cache struct {
	RedisClient *redis.Client
}

// NewCache 实例化 Cache 结构体
func NewCache() *Cache {
	return &Cache{RedisClient: Redis.GetRedis()}
}

// GetOrSet 从缓存中获取数据，如果缓存中没有数据则从dataSource获取并存入缓存
// @Param ctx context.Context 上下文
// @Param key string 缓存键
// @Param expiration time.Duration 缓存过期时间
// @Param dataSource func() (interface{}, error) 从数据源获取数据的函数
// @Return interface{} 数据
func (c *Cache) GetOrSet(ctx context.Context, key string, expiration time.Duration, dataSource func() (interface{}, error)) (interface{}, error) {
	// 从Redis中获取数据
	val, err := c.RedisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 缓存中没有数据，从dataSource获取
		data, err := dataSource()
		if err != nil {
			return nil, err
		}

		// 将数据序列化为JSON并存入缓存
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		err = c.RedisClient.Set(ctx, key, jsonData, expiration).Err()
		if err != nil {
			return nil, err
		}

		return data, nil
	} else if err != nil {
		return nil, err
	}

	// 反序列化缓存中的数据
	var result interface{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
