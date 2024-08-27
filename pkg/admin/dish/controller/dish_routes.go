package controller

import (
	"github.com/gin-gonic/gin"
	dishDao "sky-take-out-gin/pkg/admin/dish/dao"
	dishService "sky-take-out-gin/pkg/admin/dish/service"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
)

func DishRoutes(route *gin.RouterGroup) {
	// 实例化DishServiceImpl
	db := database.GetDatabaseManager()
	dao := dishDao.NewDishDaoImpl(db)
	cache := cache2.NewCache(db)
	service := dishService.NewDishServiceImpl(dao, cache)
	controller := NewDishControllerImpl(service)

	dishRoute := route.Group("/dish").Use(middleware.JWTMiddleware(middleware.Admin))
	{
		dishRoute.PUT("", controller.UpdateDish)
		dishRoute.DELETE("", controller.DeleteDish)
		dishRoute.POST("", controller.AddDish)
		dishRoute.GET("/:id", controller.SearchDishByID)
		dishRoute.GET("/list/:category_id", controller.SearchDishByCategory)
		dishRoute.GET("/page", controller.SearchDishByPage)
		dishRoute.PUT("/status/:status", controller.ChangeDishStatus)
	}
}
