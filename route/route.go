package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	adminCategoryRoute "sky-take-out-gin/pkg/admin/category/controller"
	adminDishRoute "sky-take-out-gin/pkg/admin/dish/controller"
	adminEmployeeRoute "sky-take-out-gin/pkg/admin/employee/controller"
	adminFileRoute "sky-take-out-gin/pkg/admin/file/controller"
	adminSetmealRoute "sky-take-out-gin/pkg/admin/setmeal/controller"
	adminShopRoute "sky-take-out-gin/pkg/admin/shop/controller"
	"sky-take-out-gin/pkg/common/config"
	"sky-take-out-gin/pkg/common/logger"
	"sky-take-out-gin/pkg/common/middleware"
	sseRoute "sky-take-out-gin/pkg/sse/controller"
	userAddressBookRoute "sky-take-out-gin/pkg/user/address_book/controller"
	userCategoryRoute "sky-take-out-gin/pkg/user/category/controller"
	"time"
)

// SetupHTTPRoute 用于初始化 HTTP 路由
func SetupHTTPRoute() error {
	// 初始化路由
	route := gin.New()

	route.Use(logger.GinLogger(zap.L()))
	route.Use(logger.GinRecovery(zap.L(), true))

	route.Use(middleware.TimeoutMiddleware(time.Second * 4))

	ApiV1 := route.Group("/api/v1")
	{
		userAPI := ApiV1.Group("/user")
		{
			userCategoryRoute.CategoryRoute(userAPI)
			userAddressBookRoute.AddressBookRoute(userAPI)

		}

		adminAPI := ApiV1.Group("/admin")
		{
			adminCategoryRoute.CategoryRoutes(adminAPI)
			adminEmployeeRoute.EmployeeRoutes(adminAPI)
			adminSetmealRoute.SetmealRoutes(adminAPI)
			adminShopRoute.ShopRoutes(adminAPI)
			adminFileRoute.FileRoutes(adminAPI)
			adminDishRoute.DishRoutes(adminAPI)
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
