package controller

import (
	"github.com/gin-gonic/gin"
	orderDao "sky-take-out-gin/pkg/admin/order/dao"
	orderService "sky-take-out-gin/pkg/admin/order/service"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
)

func AdminOrderRoutes(route *gin.RouterGroup) {
	// 实例化AdminOrderServiceImpl
	db := database.GetDatabaseManager()
	dao := orderDao.NewAdminOrderDaoImpl(db)
	service := orderService.NewAdminOrderService(dao)
	controller := NewAdminOrderController(service)

	orderRoute := route.Group("/order").Use(middleware.JWTMiddleware(middleware.Admin))
	{
		orderRoute.PUT("/cancel", controller.CancelOrder)
		orderRoute.GET("/statistics", controller.GetOrderStatistics)
		orderRoute.PUT("/complete/:order_id", controller.FinishOrder)
		orderRoute.POST("/rejection", controller.RejectOrder)
		orderRoute.PUT("/confirm", controller.ConfirmOrder)
		orderRoute.GET("/detail/:order_id", controller.GetOrderByID)
		orderRoute.PUT("/delivery/:order_id", controller.DeliveryOrder)
		orderRoute.GET("/condition_search", controller.ConditionSearchOrder)
	}
}
