package controller

import "github.com/gin-gonic/gin"

type UserShoppingCartControllerInterface interface {
	GetShoppingCartList(c *gin.Context)
	AddToShoppingCart(c *gin.Context)
	RemoveFromShoppingCart(c *gin.Context)
	SubOneCommodity(c *gin.Context)
}
