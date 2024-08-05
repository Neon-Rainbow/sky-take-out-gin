package controller

import "github.com/gin-gonic/gin"

type DishControllerInfertace interface {
	UpdateDish(c *gin.Context)
	DeleteDish(c *gin.Context)
	AddDish(c *gin.Context)
	SearchDishByID(c *gin.Context)
	SearchDishByCategory(c *gin.Context)
	SearchDishByPage(c *gin.Context)
	ChangeDishStatus(c *gin.Context)
}
