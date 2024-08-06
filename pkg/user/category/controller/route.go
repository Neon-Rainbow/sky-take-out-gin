package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	catogoryDao "sky-take-out-gin/pkg/user/category/dao"
	categoryService "sky-take-out-gin/pkg/user/category/service"
)

func CategoryRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := catogoryDao.NewCategoryDao(db)
	service := categoryService.NewCategoryService(dao)
	controller := NewCategoryController(service)

	categoryRoute := route.Group("/category")
	{
		categoryRoute.GET("", controller.GetCategoryList)
	}
}
