package controller

import "github.com/gin-gonic/gin"

type UserOrderControllerInterface interface {
	ReminderOrder(c *gin.Context)
	RepetitionOrder(c *gin.Context)
	GetHistoryOrder(c *gin.Context)
	CancelOrder(c *gin.Context)
	GetOrderDetail(c *gin.Context)
	SubmitOrder(c *gin.Context)
	PayOrder(c *gin.Context)
}
