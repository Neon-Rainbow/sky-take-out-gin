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
	userDishRoute "sky-take-out-gin/pkg/user/dish/controller"
	userOrderRoute "sky-take-out-gin/pkg/user/order/controller"
	userSetMealRoute "sky-take-out-gin/pkg/user/set_meal/controller"
	userShopRoute "sky-take-out-gin/pkg/user/shop/controller"
	userShoppingCartRoute "sky-take-out-gin/pkg/user/shopping_cart/controller"
	userLoginRoute "sky-take-out-gin/pkg/user/user/controller"
	"time"
)

// SetupHTTPRoute 用于初始化 HTTP 路由
func SetupHTTPRoute() error {
	// 初始化路由
	route := gin.New()

	// 日志中间件
	route.Use(logger.GinLogger(zap.L()))

	// 恢复中间件
	route.Use(logger.GinRecovery(zap.L(), true))

	// 超时中间件
	//route.Use(middleware.TimeoutMiddleware(time.Duration(config.GetConfig().Timeout) * time.Second))

	// CORS 中间件
	route.Use(middleware.CORSMiddleware())

	ApiV1 := route.Group("/api/v1")
	{
		userAPI := ApiV1.Group("/user")
		userAPI.Use(middleware.TimeoutMiddleware(time.Duration(config.GetConfig().Timeout) * time.Second))
		{
			userCategoryRoute.CategoryRoute(userAPI)
			userAddressBookRoute.AddressBookRoute(userAPI)
			userSetMealRoute.SetMealRoute(userAPI)
			userShopRoute.ShopRoute(userAPI)
			userLoginRoute.UserRoute(userAPI)
			userDishRoute.DishRoute(userAPI)
			userOrderRoute.UserOrderRoute(userAPI)
			userShoppingCartRoute.UserShoppingCartRoute(userAPI)
		}

		adminAPI := ApiV1.Group("/admin")
		adminAPI.Use(middleware.TimeoutMiddleware(time.Duration(config.GetConfig().Timeout) * time.Second))
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
