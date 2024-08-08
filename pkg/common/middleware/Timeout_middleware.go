package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// 用新的上下文替换请求的上下文
		c.Request = c.Request.WithContext(ctx)

		// 继续处理请求
		c.Next()

		// 检查上下文是否超时
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Operation timed out"})
			c.Abort()
		}
	}
}
