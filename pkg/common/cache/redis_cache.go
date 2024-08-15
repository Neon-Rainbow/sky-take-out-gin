package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"sky-take-out-gin/pkg/common/database"
	"time"
)

// Cache 结构体
type Cache struct {
	db database.DatabaseInterface
}

// NewCache 实例化 Cache 结构体
func NewCache(db database.DatabaseInterface) *Cache {
	return &Cache{db: db}
}

// GetOrSet 从缓存中获取数据，如果缓存中没有数据则从dataSource获取并存入缓存
// @Param ctx context.Context 上下文
// @Param key string 缓存键
// @Param expiration time.Duration 缓存过期时间
// @Param target interface{} 目标数据
// @Param dataSource func(ctx context.Context, args ...interface{}) (interface{}, error) 从数据源获取数据的函数
// @Param args 可变参数，用于传递给dataSource函数
// @Return interface{} 数据
func (c *Cache) GetOrSet(
	ctx context.Context,
	key string, expiration time.Duration,
	target interface{},
	dataSource func(ctx context.Context, args ...interface{}) (interface{}, error), args ...interface{},
) error {
	val, err := c.db.GetRedis().Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 缓存中没有数据，从dataSource获取
		data, err := dataSource(ctx, args...)
		if err != nil {
			return err
		}

		// 将数据复制到目标对象
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(jsonData, target)
		if err != nil {
			return err
		}

		// 启动一个goroutine将数据保存到Redis中
		go func() {
			// 将数据序列化为JSON并保存到Redis
			err := c.db.GetRedis().Set(ctx, key, jsonData, expiration).Err()
			if err != nil {
				// 处理Redis存储错误，例如记录日志
				return
			}
		}()

		// 立即返回获取的数据
		return nil
	} else if err != nil {
		return err
	}

	// 反序列化缓存中的数据到目标结构
	err = json.Unmarshal([]byte(val), target)
	if err != nil {
		return err
	}

	return nil
}

// Invalidate 删除 Redis 中的缓存数据
// @Param ctx context.Context 上下文
// @Param key string 缓存键
// @Return error 错误信息
func (c *Cache) Invalidate(ctx context.Context, key string) error {
	// 尝试删除 Redis 中的缓存数据
	err := c.db.GetRedis().Del(ctx, key).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}

// InvalidatePattern 删除 Redis 中符合模式的缓存数据
// @Param ctx context.Context 上下文
// @Param pattern string 模式
// @Return error 错误信息
// 例如：InvalidatePattern(ctx, "user:*") 删除所有以 user: 开头的键
func (c *Cache) InvalidatePattern(ctx context.Context, pattern string) error {
	// 获取所有符合模式的键
	keys, err := c.db.GetRedis().Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	// 删除所有符合模式的键
	err = c.db.GetRedis().Del(ctx, keys...).Err()
	if err != nil {
		return err
	}

	return nil
}

// InvalidateAll 删除 Redis 中所有缓存数据
// @Param ctx context.Context 上下文
// @Return error 错误信息
func (c *Cache) InvalidateAll(ctx context.Context) error {
	// 删除 Redis 中所有缓存数据
	err := c.db.GetRedis().FlushDB(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
