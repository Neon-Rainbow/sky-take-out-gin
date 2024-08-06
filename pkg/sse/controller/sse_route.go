package sseRoute

import "github.com/gin-gonic/gin"

func SSERoute(route *gin.RouterGroup) {
	sse_route := route

	sse_route.GET("/", SSEHandler)
	sse_route.POST("/send_message", SendMessageHandler)
}
