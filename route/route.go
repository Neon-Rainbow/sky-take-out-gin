package route

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/route/api/v1/admin"
)

// SetupHTTPRoute 用于初始化 HTTP 路由
func SetupHTTPRoute() *gin.Engine {
	route := gin.Default()
	ApiV1 := route.Group("/api/v1")
	{
		UserApi := ApiV1.Group("/user")
		{
			admin.CategoryRoutes(UserApi)
		}
		AdminApi := ApiV1.Group("/admin")
		{

		}
	}
	return route
}
