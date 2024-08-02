package route

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/route/api/v1/admin"
)

// SetupHTTPRoute 用于初始化 HTTP 路由
func SetupHTTPRoute() error {
	route := gin.Default()
	ApiV1 := route.Group("/api/v1")
	{
		_ = ApiV1.Group("/user")
		{

		}
		AdminApi := ApiV1.Group("/admin")
		{
			admin.CategoryRoutes(AdminApi)
			admin.EmployeeRoutes(AdminApi)
		}
	}
	err := route.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
