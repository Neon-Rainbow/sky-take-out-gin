package controller

import (
	"github.com/gin-gonic/gin"
	daoCategory "sky-take-out-gin/pkg/admin/category/dao"
	serviceCategory "sky-take-out-gin/pkg/admin/category/service"
	"sky-take-out-gin/pkg/common/database"
)

func CategoryRoutes(route *gin.RouterGroup) {
	// 实例化CategoryDaoImpl

	db := database.GetDatabaseManager()
	dao := daoCategory.NewCategoryDaoImpl(db)
	service := serviceCategory.NewCategoryService(dao)
	controller := NewAdminCategoryControllerImpl(service)

	categoryRoute := route.Group("/category")
	{
		categoryRoute.PUT("/", controller.UpdateCategory)
		categoryRoute.GET("/page", controller.GetCategoryPage)
		categoryRoute.POST("/status/:status", controller.ChangeCategoryStatus)
		categoryRoute.POST("/", controller.CreateCategory)
		categoryRoute.DELETE("/", controller.DeleteCategory)
		categoryRoute.GET("/list", controller.GetCategoryListByType)
	}
}
