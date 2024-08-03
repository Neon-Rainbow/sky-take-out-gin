package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/code"
	"sky-take-out-gin/internal/controller"
	"sky-take-out-gin/utils/jwt"
	"strings"
)

// JWTMiddleware JWT中间件
func JWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Authorization header 不能为空")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
			controller.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Authorization header 的格式必须为 Bearer {token}")

			c.Abort()
			return
		}

		tkn := parts[1]
		myClaims, err := jwt.ParseToken(tkn)
		if err != nil {
			controller.ResponseErrorWithMsg(c, http.StatusUnauthorized, code.RequestUnauthorized, "Token 无效")
			//httpController.ResponseErrorWithMessage(c, code.RequestUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if myClaims.TokenType != "access" {
			controller.ResponseErrorWithMsg(
				c,
				http.StatusUnauthorized,
				code.RequestUnauthorized,
				fmt.Sprintf("无效的token格式,应该为access token,实际为%s", myClaims.TokenType))
			c.Abort()
			return
		}

		c.Set("userID", myClaims.UserID)
		c.Set("username", myClaims.Username)
		c.Next()
	}
}
