package admin

import (
	"github.com/gin-gonic/gin"
	setmealController "sky-take-out-gin/internal/controller/admin/setmeal"
	setmealDao "sky-take-out-gin/internal/dao/admin/setmeal"
	setmealService "sky-take-out-gin/internal/service/admin/setmeal"
	"sky-take-out-gin/middleware"
)

func SetmealRoutes(route *gin.RouterGroup) {
	dao := setmealDao.NewSetmealDAOImpl()
	service := setmealService.NewSetmealServiceImpl(dao)
	controller := setmealController.NewSetmealControllerImpl(service)

	setmeal_route := route.Group("/setmeal").Use(middleware.JWTMiddleware())
	{
		setmeal_route.PUT("/", controller.UpdateSetmeal)
		setmeal_route.GET("/page", controller.GetSetmealPage)
		setmeal_route.POST("/status/:status", controller.ChangeSetmealStatus)
		setmeal_route.DELETE("/", controller.DeleteSetmeals)
		setmeal_route.POST("/", controller.CreateSetmeals)
		setmeal_route.GET("/:id", controller.GetSetmealsByID)
	}
}
