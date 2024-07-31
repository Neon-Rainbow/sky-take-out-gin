package route

import "github.com/gin-gonic/gin"

// SetupHTTPRoute 用于初始化 HTTP 路由
func SetupHTTPRoute() *gin.Engine {
	route := gin.Default()
	ApiV1 := route.Group("/api/v1")
	{
		UserApi := ApiV1.Group("/user")
		{

		}
		AdminApi := ApiV1.Group("/admin")
		{

		}
	}
}
