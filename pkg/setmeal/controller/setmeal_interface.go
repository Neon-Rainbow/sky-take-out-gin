package setmeal

import "github.com/gin-gonic/gin"

type SetmealControllerInterface interface {
	UpdateSetmeal(c *gin.Context)
	GetSetmealPage(c *gin.Context)
	ChangeSetmealStatus(c *gin.Context)
	DeleteSetmeals(c *gin.Context)
	CreateSetmeals(c *gin.Context)
	GetSetmealsByID(c *gin.Context)
}
