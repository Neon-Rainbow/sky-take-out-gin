package controller

import "github.com/gin-gonic/gin"

type ShopStatusControllerInterface interface {
	// GetShopStatus 获取店铺状态
	GetShopStatus(c *gin.Context)
}
