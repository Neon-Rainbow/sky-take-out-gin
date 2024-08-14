package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	setMealDao "sky-take-out-gin/pkg/user/set_meal/dao"
	setMealService "sky-take-out-gin/pkg/user/set_meal/service"
)

// SetMealRoute 设置套餐路由
func SetMealRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := setMealDao.NewSetMealDaoImpl(db)
	service := setMealService.NewSetMealServiceImpl(dao)
	controller := NewSetMealController(service)

	setMealRoute := route.Group("/setmeal")
	{
		setMealRoute.GET("/list", controller.GetSetMealList)
		setMealRoute.GET("/dish/:id", controller.GetSetMealDetail)
	}
}
