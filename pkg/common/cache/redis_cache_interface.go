package cache

import (
	"context"
	"time"
)

type RedisCacheInterface interface {
	// GetOrSet 从缓存中获取数据，如果缓存中没有数据则从dataSource获取并存入缓存
	GetOrSet(ctx context.Context, key string, expiration time.Duration, target interface{}, dataSource func(ctx context.Context, args ...interface{}) (interface{}, error), args ...interface{}) error

	// Invalidate 删除 Redis 中的缓存数据
	Invalidate(ctx context.Context, key string) error

	// InvalidatePattern 删除 Redis 中的缓存数据
	InvalidatePattern(ctx context.Context, pattern string) error

	// InvalidateAll 删除 Redis 中的所有缓存数据
	InvalidateAll(ctx context.Context) error
}
