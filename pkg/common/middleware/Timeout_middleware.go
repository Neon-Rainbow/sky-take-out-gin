package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeoutMiddleware 用于设置请求超时时间
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建一个带有超时的上下文
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// 用新的上下文替换请求的上下文
		c.Request = c.Request.WithContext(ctx)

		// 创建一个通道，用于接收处理完成信号
		done := make(chan struct{})

		// 使用协程执行 c.Next() 并发送完成信号
		go func() {
			c.Next()
			done <- struct{}{}
		}()

		// 等待处理完成或上下文超时
		select {
		case <-ctx.Done():
			// 如果上下文在操作完成前超时
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				c.JSON(http.StatusGatewayTimeout, gin.H{
					"message": "请求超时",
				})
				c.Abort()
			}
		case <-done:
			// 处理完成
		}
	}
}
