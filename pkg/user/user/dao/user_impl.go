package dao

import (
	"context"
	"errors"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
	"strconv"
	"time"
)

type UserTokenDaoImpl struct {
	db database.DatabaseInterface
}

func (dao *UserTokenDaoImpl) SaveToken(ctx context.Context, userID uint, accessToken string, refreshToken string, expiration time.Duration) error {
	userKey := generateUserKey(userID)

	// 将用户的访问和刷新 Token 存储到 Redis 中，并设置过期时间
	err := dao.db.GetRedis().HSet(ctx, userKey, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}).Err()
	if err != nil {
		return err
	}

	// 设置过期时间
	return dao.db.GetRedis().Expire(ctx, userKey, expiration).Err()
}

func (dao *UserTokenDaoImpl) GetTokens(ctx context.Context, userID uint) (accessToken string, refreshToken string, error error) {
	userKey := generateUserKey(userID)
	tokens, err := dao.db.GetRedis().HMGet(ctx, userKey, "access_token", "refresh_token").Result()
	if err != nil {
		return "", "", err
	}
	if len(tokens) == 0 {
		return "", "", errors.New("tokens not found")
	}

	return tokens[0].(string), tokens[1].(string), nil
}

func (dao *UserTokenDaoImpl) DeleteTokens(ctx context.Context, userID uint) error {
	userKey := generateUserKey(userID)
	return dao.db.GetRedis().Del(ctx, userKey).Err()
}

// ValidateAccessToken 验证 access token 是否有效
func (dao *UserTokenDaoImpl) ValidateAccessToken(ctx context.Context, userID uint, accessToken string) (bool, error) {
	// 验证 access token 的逻辑
	panic("implement me")
}

// ValidateRefreshToken 验证 refresh token 是否有效
func (dao *UserTokenDaoImpl) ValidateRefreshToken(ctx context.Context, userID uint, refreshToken string) (bool, error) {
	// 验证 refresh token 的逻辑
	token, err := dao.db.GetRedis().HMGet(ctx, generateUserKey(userID), "refresh_token").Result()
	if err != nil {
		return false, err
	}

	return token[0].(string) == refreshToken, nil
}

type UserDaoImpl struct {
	db database.DatabaseInterface
}

func (dao *UserDaoImpl) GetUserByUsername(ctx context.Context, username string) (userID uint, hashedPassword string, error error) {
	var user model.User
	err := dao.db.GetDB().Where("username = ?", username).First(&user).Error
	if err != nil {
		return 0, "", err
	}
	return user.ID, user.Password, nil
}

// CreateUser 创建用户
func (dao *UserDaoImpl) CreateUser(ctx context.Context, username string, hashedPassword string) (userID uint, error error) {
	user := model.User{Username: username, Password: hashedPassword}
	err := dao.db.GetDB().Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func NewUserDaoImpl(db database.DatabaseInterface) *UserDaoImpl {
	return &UserDaoImpl{db: db}
}

func NewUserTokenDaoImpl(db database.DatabaseInterface) *UserTokenDaoImpl {
	return &UserTokenDaoImpl{db: db}
}

// generateUserKey 生成用户在 Redis 中的 Key
func generateUserKey(userID uint) string {
	return "user:" + strconv.Itoa(int(userID))
}
