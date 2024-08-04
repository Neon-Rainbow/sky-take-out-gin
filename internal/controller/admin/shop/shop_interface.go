package shop

import "github.com/gin-gonic/gin"

type ShopControllerInterface interface {
	GetShopStatus(c *gin.Context)
	SetShopStatus(c *gin.Context)
}
