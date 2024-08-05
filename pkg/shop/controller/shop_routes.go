package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/middleware"
	shopDao "sky-take-out-gin/pkg/shop/dao"
	shopService "sky-take-out-gin/pkg/shop/service"
)

// ShopRoutes 商店信息相关路由
func ShopRoutes(route *gin.RouterGroup) {
	dao := shopDao.NewShopDaoImpl()
	service := shopService.NewShopServiceImpl(dao)
	controller := NewShopControllerImpl(service)

	shopRoute := route.Group("/shop").Use(middleware.JWTMiddleware())
	{
		shopRoute.GET("/status", controller.GetShopStatus)
		shopRoute.PUT("/status/:status", controller.SetShopStatus)
	}
}
