package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	dishDao "sky-take-out-gin/pkg/dish/dao"
	dishService "sky-take-out-gin/pkg/dish/service"
)

func DishRoutes(route *gin.RouterGroup) {
	// 实例化DishServiceImpl
	db := database.GetDatabaseManager()
	dao := dishDao.NewDishDaoImpl(db)
	service := dishService.NewDishServiceImpl(dao)
	controller := NewDishControllerImpl(service)

	dishRoute := route.Group("/dish")
	{
		dishRoute.PUT("/", controller.UpdateDish)
		dishRoute.DELETE("/", controller.DeleteDish)
		dishRoute.POST("/", controller.AddDish)
		dishRoute.GET("/:id", controller.SearchDishByID)
		dishRoute.GET("/list/:category_id", controller.SearchDishByCategory)
		dishRoute.GET("/page", controller.SearchDishByPage)
		dishRoute.PUT("/status/:status", controller.ChangeDishStatus)
	}
}
