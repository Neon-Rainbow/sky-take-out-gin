package controller

import "github.com/gin-gonic/gin"

type DishControllerInterface interface {
	GetDishByID(ctx *gin.Context)
}
