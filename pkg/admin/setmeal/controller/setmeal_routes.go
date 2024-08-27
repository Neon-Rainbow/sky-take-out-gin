package controller

import (
	"github.com/gin-gonic/gin"
	setmealDao "sky-take-out-gin/pkg/admin/setmeal/dao"
	setmealService "sky-take-out-gin/pkg/admin/setmeal/service"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
)

func SetMealRoutes(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := setmealDao.NewSetmealDAOImpl(db)
	cache := cache2.NewCache(db)
	service := setmealService.NewSetmealServiceImpl(dao, cache)
	controller := NewSetmealControllerImpl(service)

	setMealRoute := route.Group("/set_meal").Use(middleware.JWTMiddleware(middleware.Admin))
	{
		setMealRoute.PUT("/", controller.UpdateSetmeal)
		setMealRoute.GET("/page", controller.GetSetmealPage)
		setMealRoute.POST("/status/:status", controller.ChangeSetmealStatus)
		setMealRoute.DELETE("/", controller.DeleteSetmeals)
		setMealRoute.POST("/", controller.CreateSetmeals)
		setMealRoute.GET("/:id", controller.GetSetmealsByID)
	}
}
