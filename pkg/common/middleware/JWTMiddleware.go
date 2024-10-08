package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/JWT"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/response"
	dao2 "sky-take-out-gin/pkg/user/user/dao"
	"strings"
)

const (
	// Admin 管理员
	Admin = "employee"

	// User 用户
	User = "user"
)

// JWTMiddleware JWT中间件
func JWTMiddleware(userType string) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Authorization header 不能为空")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
			response.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Authorization header 的格式必须为 Bearer {token}")

			c.Abort()
			return
		}

		tkn := parts[1]
		myClaims, err := JWT.ParseToken(tkn)
		if err != nil {
			response.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Token 无效")
			//httpController.ResponseErrorWithMessage(c, code.RequestUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if myClaims.TokenType != "access" {
			response.ResponseErrorWithMsg(
				c,
				http.StatusUnauthorized,
				code.RequestUnauthorized,
				fmt.Sprintf("无效的token格式,应该为access token,实际为%s", myClaims.TokenType))
			c.Abort()
			return
		}

		if myClaims.UserType != userType {
			response.ResponseErrorWithMsg(
				c,
				http.StatusForbidden,
				code.RequestForbidden,
				fmt.Sprintf("无效的用户类型,需要%s, 实际为%s", userType, myClaims.UserType))
			c.Abort()
			return
		}

		if myClaims.UserType == User {
			db := database.GetDatabaseManager()
			dao := dao2.NewUserTokenDaoImpl(db)
			var redisAccessToken string
			redisAccessToken, _, err = dao.GetTokens(context.Background(), myClaims.UserID)
			if err != nil {
				response.ResponseErrorWithMsg(c, http.StatusInternalServerError, code.ServerError, "查询Redis中的token失败")
				c.Abort()
				return
			}
			if redisAccessToken != tkn {
				response.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Token 已过期")
				c.Abort()
				return
			}
		}

		c.Set("userID", myClaims.UserID)
		c.Set("username", myClaims.Username)
		c.Set("userType", myClaims.UserType)
		c.Next()
	}
}
