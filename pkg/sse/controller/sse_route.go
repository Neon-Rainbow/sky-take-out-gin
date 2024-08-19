package sseRoute

import "github.com/gin-gonic/gin"

func SSERoute(route *gin.RouterGroup) {
	sseRoute := route

	sseRoute.GET("", SSEHandler)
	sseRoute.POST("/send_message", SendMessageHandler)
}
