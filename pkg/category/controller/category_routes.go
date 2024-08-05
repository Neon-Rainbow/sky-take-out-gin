package admin

import (
	"github.com/gin-gonic/gin"
	controllerCategory "sky-take-out-gin/internal/controller/admin/category"
	daoCategory "sky-take-out-gin/internal/dao/admin/category"
	serviceCategory "sky-take-out-gin/internal/service/admin/category"
)

func CategoryRoutes(route *gin.RouterGroup) {
	// 实例化CategoryDaoImpl
	dao := daoCategory.NewCategoryDaoImpl()
	service := serviceCategory.NewCategoryService(dao)
	controller := controllerCategory.NewAdminCategoryControllerImpl(service)

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
