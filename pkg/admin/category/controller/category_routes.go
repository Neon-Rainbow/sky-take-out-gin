package controller

import (
	"github.com/gin-gonic/gin"
	daoCategory "sky-take-out-gin/pkg/admin/category/dao"
	serviceCategory "sky-take-out-gin/pkg/admin/category/service"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
)

func CategoryRoutes(route *gin.RouterGroup) {
	// 实例化CategoryDaoImpl

	db := database.GetDatabaseManager()
	dao := daoCategory.NewCategoryDaoImpl(db)
	cache := cache2.NewCache(db)
	service := serviceCategory.NewCategoryService(dao, cache)
	controller := NewAdminCategoryControllerImpl(service)

	categoryRoute := route.Group("/category").Use(middleware.JWTMiddleware(middleware.Admin))
	{
		categoryRoute.PUT("/", controller.UpdateCategory)
		categoryRoute.GET("/page", controller.GetCategoryPage)
		categoryRoute.POST("/status/:status", controller.ChangeCategoryStatus)
		categoryRoute.POST("/", controller.CreateCategory)
		categoryRoute.DELETE("/", controller.DeleteCategory)
		categoryRoute.GET("/list", controller.GetCategoryListByType)
	}
}
