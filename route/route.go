package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	categoryRoute "sky-take-out-gin/pkg/category/controller"
	"sky-take-out-gin/pkg/common/config"
	"sky-take-out-gin/pkg/common/logger"
	dishRoute "sky-take-out-gin/pkg/dish/controller"
	employeeRoute "sky-take-out-gin/pkg/employee/controller"
	fileRoute "sky-take-out-gin/pkg/file/controller"
	setmealRoute "sky-take-out-gin/pkg/setmeal/controller"
	shopRoute "sky-take-out-gin/pkg/shop/controller"
	sseRoute "sky-take-out-gin/pkg/sse/controller"
)

// SetupHTTPRoute 用于初始化 HTTP 路由
func SetupHTTPRoute() error {
	// 初始化路由
	route := gin.New()

	route.Use(logger.GinLogger(zap.L()))
	route.Use(logger.GinRecovery(zap.L(), true))

	ApiV1 := route.Group("/api/v1")
	{
		_ = ApiV1.Group("/user")
		{

		}
		AdminApi := ApiV1.Group("/admin")
		{
			categoryRoute.CategoryRoutes(AdminApi)
			employeeRoute.EmployeeRoutes(AdminApi)
			setmealRoute.SetmealRoutes(AdminApi)
			shopRoute.ShopRoutes(AdminApi)
			fileRoute.FileRoutes(AdminApi)
			dishRoute.DishRoutes(AdminApi)
		}

		sseApi := ApiV1.Group("/sse")
		{
			sseRoute.SSERoute(sseApi)
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
