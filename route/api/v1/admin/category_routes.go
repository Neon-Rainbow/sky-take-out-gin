package admin

import "github.com/gin-gonic/gin"

func CategoryRoutes(route *gin.RouterGroup) {
	category_route := route.Group("/category")
	{
		category_route.PUT("/")
	}
}
