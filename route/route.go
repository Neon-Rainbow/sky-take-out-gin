package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/config"
	adminRoute "sky-take-out-gin/route/api/v1/admin"
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
			adminRoute.CategoryRoutes(AdminApi)
			adminRoute.EmployeeRoutes(AdminApi)
			adminRoute.SetmealRoutes(AdminApi)
			adminRoute.ShopRoutes(AdminApi)
			adminRoute.FileRoutes(AdminApi)
			adminRoute.DishRoutes(AdminApi)
		}
	}
	err := route.Run(fmt.Sprintf(
		"%s:%d",
		config.GetConfig().ServerConfig.Host,
		config.GetConfig().ServerConfig.Port),
	)
	if err != nil {
		return err
	}
	return nil
}
