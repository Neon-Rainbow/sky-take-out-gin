package controller

import (
	"github.com/gin-gonic/gin"
	shopDao "sky-take-out-gin/pkg/admin/shop/dao"
	shopService "sky-take-out-gin/pkg/admin/shop/service"
	"sky-take-out-gin/pkg/common/database"
)

// ShopRoutes 商店信息相关路由
func ShopRoutes(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := shopDao.NewShopDaoImpl(db)
	service := shopService.NewShopServiceImpl(dao)
	controller := NewShopControllerImpl(service)

	shopRoute := route.Group("/shop")
	{
		shopRoute.GET("/status", controller.GetShopStatus)
		shopRoute.PUT("/status/:status", controller.SetShopStatus)
	}
}
