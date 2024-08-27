package controller

import "github.com/gin-gonic/gin"

type AdminOrderControllerInterface interface {
	CancelOrder(c *gin.Context)
	GetOrderStatistics(c *gin.Context)
	FinishOrder(c *gin.Context)
	RejectOrder(c *gin.Context)
	ConfirmOrder(c *gin.Context)
	GetOrderByID(c *gin.Context)
	DeliveryOrder(c *gin.Context)
	ConditionSearchOrder(c *gin.Context)
}
