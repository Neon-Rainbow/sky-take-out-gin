package JWT

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"sky-take-out-gin/pkg/common/config"
	"time"
)

const (
	// accessTokenExpireDuration 是访问令牌的过期时间
	accessTokenExpireDuration = time.Hour * 24
	// refreshTokenExpireDuration 是刷新令牌的过期时间
	refreshTokenExpireDuration = time.Hour * 24 * 15
)

type MyClaims struct {
	Username  string `json:"username"`
	UserID    int64  `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// GenerateToken generates both access and refresh JWT tokens for a given user
// @title GenerateToken
// @description Generates JWT tokens
// @param username string The username of the user
// @param userId uint The user ID
// @return accessToken string The access token
// @return refreshToken string The refresh token
// @return err error information
func GenerateToken(username string, userId int64, isAdmin bool) (accessToken string, refreshToken string, err error) {
	// Retrieve the secret key from the configuration
	var jwtSecret = config.GetConfig().SecretConfig.JWTSecret
	var mySecret = []byte(jwtSecret)

	// Channels to handle success and error signals
	successChannel := make(chan bool, 2)
	errorChannel := make(chan error, 2)

	// Create a context with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// function 用于生成JWT令牌
	// @Param username string 用户名
	// @Param userId int64 用户ID
	// @Param tokenType string 令牌类型
	// @Param expireDuration time.Duration 令牌过期时间
	// @Return string 令牌字符串
	// @Return error 错误信息
	var function = func(username string, userId int64, tokenType string, expireDuration time.Duration) (string, error) {
		claims := MyClaims{
			Username:  username,
			UserID:    userId,
			TokenType: tokenType,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	// 使用两个goroutine分别生成访问令牌和刷新令牌
	go func() {
		accessToken, err = function(username, int64(userId), "access", accessTokenExpireDuration)
		if err != nil {
			errorChannel <- err
			return
		}
		successChannel <- true
	}()

	go func() {
		refreshToken, err = function(username, int64(userId), "refresh", refreshTokenExpireDuration)
		if err != nil {
			errorChannel <- err
			return
		}
		successChannel <- true
	}()

	// Loop to wait for both tokens to be generated or an error/timeout to occur
	for i := 0; i < 2; i++ {
		select {
		case <-successChannel:
			// Successfully generated a token
		case err = <-errorChannel:
			// An error occurred while generating a token
			return "", "", err
		case <-ctx.Done():
			// Context timeout
			return "", "", ctx.Err()
		}
	}

	// Return the generated access and refresh tokens
	return accessToken, refreshToken, nil
}

// ParseToken 用于解析JWT令牌
// @Param tokenString string JWT令牌字符串
// @Return *MyClaims JWT令牌的声明
func ParseToken(tokenString string) (*MyClaims, error) {
	// Retrieve the secret key from the configuration
	var jwtSecret = config.GetConfig().SecretConfig.JWTSecret
	var mySecret = []byte(jwtSecret) // Custom secret key

	// Parse the token with the custom claims structure
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		})
	if err != nil {
		return nil, err
	}

	// Validate the token and extract the claims
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
