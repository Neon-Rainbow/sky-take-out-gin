package controller

import (
	"github.com/gin-gonic/gin"
	daoCategory "sky-take-out-gin/pkg/category/dao"
	serviceCategory "sky-take-out-gin/pkg/category/service"
)

func CategoryRoutes(route *gin.RouterGroup) {
	// 实例化CategoryDaoImpl
	dao := daoCategory.NewCategoryDaoImpl()
	service := serviceCategory.NewCategoryService(dao)
	controller := NewAdminCategoryControllerImpl(service)

	category_route := route.Group("/category")
	{
		category_route.PUT("/", controller.UpdateCategory)
		category_route.GET("/page", controller.GetCategoryPage)
		category_route.POST("/status/:status", controller.ChangeCategoryStatus)
		category_route.POST("/", controller.CreateCategory)
		category_route.DELETE("/", controller.DeleteCategory)
		category_route.GET("/list", controller.GetCategoryListByType)
	}
}
