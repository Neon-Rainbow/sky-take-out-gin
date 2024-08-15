package controller

import (
	"github.com/gin-gonic/gin"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	catogoryDao "sky-take-out-gin/pkg/user/category/dao"
	categoryService "sky-take-out-gin/pkg/user/category/service"
)

func CategoryRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := catogoryDao.NewCategoryDao(db)
	cache := cache2.NewCache(db)
	service := categoryService.NewCategoryService(dao, cache)
	controller := NewCategoryController(service)

	categoryRoute := route.Group("/category")
	{
		categoryRoute.GET("", controller.GetCategoryList)
	}
}
