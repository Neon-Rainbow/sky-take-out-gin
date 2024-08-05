package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/middleware"
	setmealDao "sky-take-out-gin/pkg/setmeal/dao"
	setmealService "sky-take-out-gin/pkg/setmeal/service"
)

func SetmealRoutes(route *gin.RouterGroup) {
	dao := setmealDao.NewSetmealDAOImpl()
	service := setmealService.NewSetmealServiceImpl(dao)
	controller := NewSetmealControllerImpl(service)

	setmealRoute := route.Group("/setmeal").Use(middleware.JWTMiddleware())
	{
		setmealRoute.PUT("/", controller.UpdateSetmeal)
		setmealRoute.GET("/page", controller.GetSetmealPage)
		setmealRoute.POST("/status/:status", controller.ChangeSetmealStatus)
		setmealRoute.DELETE("/", controller.DeleteSetmeals)
		setmealRoute.POST("/", controller.CreateSetmeals)
		setmealRoute.GET("/:id", controller.GetSetmealsByID)
	}
}
