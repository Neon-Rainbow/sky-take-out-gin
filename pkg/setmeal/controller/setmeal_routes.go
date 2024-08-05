package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/middleware"
	setmealController "sky-take-out-gin/pkg/setmeal/controller"
	setmealDao "sky-take-out-gin/pkg/setmeal/dao"
	setmealService "sky-take-out-gin/pkg/setmeal/service"
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
