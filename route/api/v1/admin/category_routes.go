package admin

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/internal/controller/admin/category"
)

func CategoryRoutes(route *gin.RouterGroup) {
	//TODO:controller应该不在这里声明,之后需要修改
	controller := category.NewAdminCategoryControllerImpl()
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
