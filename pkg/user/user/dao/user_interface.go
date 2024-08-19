package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"time"
)

// UserTokenDAOInterface 定义了与 Redis 交互以管理用户 Token 的操作
type UserTokenDAOInterface interface {
	// SaveToken 将用户的访问和刷新 Token 存储到 Redis 中，并设置过期时间
	SaveToken(ctx context.Context, userID uint, accessToken string, refreshToken string, expiration time.Duration) error

	// GetTokens 根据用户 ID 从 Redis 中获取存储的访问和刷新 Token
	GetTokens(ctx context.Context, userID uint) (accessToken string, refreshToken string, error error)

	// DeleteTokens 从 Redis 中删除指定用户的访问和刷新 Token，用于登出操作
	DeleteTokens(ctx context.Context, userID uint) error

	// ValidateToken 检查 Redis 中是否存在指定的访问 Token，用于验证 Token 是否有效
	//ValidateToken(ctx context.Context, userID uint, accessToken string, refreshToken string) (bool, error)

	ValidateAccessToken(ctx context.Context, userID uint, accessToken string) (bool, error)

	ValidateRefreshToken(ctx context.Context, userID uint, refreshToken string) (bool, error)
}

// UserDaoInterface 定义了与数据库交互以管理用户信息的操作
type UserDaoInterface interface {
	// GetUserByUsername 根据用户名从数据库中检索用户信息，包括用户的 ID 和加密后的密码
	GetUserByUsername(ctx context.Context, username string) (userID uint, hashedPassword string, error error)
	CreateUser(ctx context.Context, username string, hashedPassword string) (userID uint, error error)
	GetUserByID(ctx context.Context, userID uint) (user *model.User, error error)
}
