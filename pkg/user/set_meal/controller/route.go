package controller

import (
	"github.com/gin-gonic/gin"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
	setMealDao "sky-take-out-gin/pkg/user/set_meal/dao"
	setMealService "sky-take-out-gin/pkg/user/set_meal/service"
)

// SetMealRoute 设置套餐路由
func SetMealRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := setMealDao.NewSetMealDaoImpl(db)
	cache := cache2.NewCache(db)
	service := setMealService.NewSetMealServiceImpl(dao, cache)
	controller := NewSetMealController(service)

	setMealRoute := route.Group("/set_meal").Use(middleware.JWTMiddleware(middleware.User))
	{
		setMealRoute.GET("/list", controller.GetSetMealList)
		setMealRoute.GET("/dish/:id", controller.GetSetMealDetail)
	}
}
