package admin

import (
	"github.com/gin-gonic/gin"
	shopController "sky-take-out-gin/internal/controller/admin/shop"
	shopDao "sky-take-out-gin/internal/dao/admin/shop"
	shopService "sky-take-out-gin/internal/service/admin/shop"
	"sky-take-out-gin/middleware"
)

// ShopRoutes 商店信息相关路由
func ShopRoutes(route *gin.RouterGroup) {
	dao := shopDao.NewShopDaoImpl()
	service := shopService.NewShopServiceImpl(dao)
	controller := shopController.NewShopControllerImpl(service)

	shop_route := route.Group("/shop").Use(middleware.JWTMiddleware())
	{
		shop_route.GET("/status", controller.GetShopStatus)
		shop_route.PUT("/status/:status", controller.SetShopStatus)
	}
}
