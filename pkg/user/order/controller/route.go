package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	dao2 "sky-take-out-gin/pkg/user/address_book/dao"
	"sky-take-out-gin/pkg/user/order/dao"
	"sky-take-out-gin/pkg/user/order/service"
	dao4 "sky-take-out-gin/pkg/user/shopping_cart/dao"
	dao3 "sky-take-out-gin/pkg/user/user/dao"
)

func UserOrderRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	userOrderDao := dao.NewUserOrderDaoImpl(db)
	userAddressDao := dao2.NewAddressBookDaoImpl(db)
	userInfoDao := dao3.NewUserDaoImpl(db)
	shoppingCartDao := dao4.NewUserShoppingCartRedisImpl(db)
	userOrderService := service.NewUserOrderServiceImpl(userOrderDao, userAddressDao, userInfoDao, shoppingCartDao)
	userOrderController := NewUserOrderController(userOrderService)

	orderRoute := route.Group("/order")
	{
		orderRoute.GET("/reminder/:orderID", userOrderController.ReminderOrder)
		orderRoute.POST("/repetition/:orderID", userOrderController.RepetitionOrder)
		orderRoute.GET("/history", userOrderController.GetHistoryOrder)
		orderRoute.PUT("/cancel/:orderID", userOrderController.CancelOrder)
		orderRoute.GET("/detail/:orderID", userOrderController.GetOrderDetail)
		orderRoute.POST("/submit", userOrderController.SubmitOrder)
		orderRoute.PUT("/pay", userOrderController.PayOrder)
	}
}
