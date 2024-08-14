package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	shopDao "sky-take-out-gin/pkg/user/shop/dao"
	shopService "sky-take-out-gin/pkg/user/shop/service"
)

// ShopRoute 商店状态路由
func ShopRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := shopDao.NewShopDaoImpl(db)
	service := shopService.NewShopStatusServiceImpl(dao)
	controller := NewShopStatusControllerImpl(service)

	shopRoute := route.Group("/shop")
	{
		shopRoute.GET("/status", controller.GetShopStatus)
	}
}
