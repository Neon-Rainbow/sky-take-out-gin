package controller

import (
	"github.com/gin-gonic/gin"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/user/dish/dao"
	"sky-take-out-gin/pkg/user/dish/service"
)

func DishRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dishDao := dao.NewDishDao(db)
	cache := cache2.NewCache(db)
	dishService := service.NewDishService(dishDao, cache)
	controller := NewDishController(dishService)

	route.GET("/dish/list", controller.GetDishByID)
}
