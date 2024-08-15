package controller

import (
	"github.com/gin-gonic/gin"
	setmealDao "sky-take-out-gin/pkg/admin/setmeal/dao"
	setmealService "sky-take-out-gin/pkg/admin/setmeal/service"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
)

func SetmealRoutes(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := setmealDao.NewSetmealDAOImpl(db)
	cache := cache2.NewCache(db)
	service := setmealService.NewSetmealServiceImpl(dao, cache)
	controller := NewSetmealControllerImpl(service)

	setmealRoute := route.Group("/setmeal")
	{
		setmealRoute.PUT("/", controller.UpdateSetmeal)
		setmealRoute.GET("/page", controller.GetSetmealPage)
		setmealRoute.POST("/status/:status", controller.ChangeSetmealStatus)
		setmealRoute.DELETE("/", controller.DeleteSetmeals)
		setmealRoute.POST("/", controller.CreateSetmeals)
		setmealRoute.GET("/:id", controller.GetSetmealsByID)
	}
}
