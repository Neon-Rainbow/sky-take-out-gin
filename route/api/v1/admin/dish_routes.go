package admin

import (
	"github.com/gin-gonic/gin"
	dishController "sky-take-out-gin/internal/controller/admin/dish"
	dishDao "sky-take-out-gin/internal/dao/admin/dish"
	dishService "sky-take-out-gin/internal/service/admin/dish"
)

func DishRoutes(route *gin.RouterGroup) {
	// 实例化DishServiceImpl
	dao := dishDao.NewDishDaoImpl()
	service := dishService.NewDishServiceImpl(dao)
	controller := dishController.NewDishControllerImpl(service)

	dish_route := route.Group("/dish")
	{
		dish_route.PUT("/", controller.UpdateDish)
		dish_route.DELETE("/", controller.DeleteDish)
		dish_route.POST("/", controller.AddDish)
		dish_route.GET("/:id", controller.SearchDishByID)
		dish_route.GET("/list/:category_id", controller.SearchDishByCategory)
		dish_route.GET("/page", controller.SearchDishByPage)
		dish_route.PUT("/status/:status", controller.ChangeDishStatus)
	}
}
