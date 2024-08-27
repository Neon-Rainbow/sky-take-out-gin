package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
	dao2 "sky-take-out-gin/pkg/user/user/dao"
	service2 "sky-take-out-gin/pkg/user/user/service"
)

// UserRoute 是用户模块的路由
func UserRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := dao2.NewUserDaoImpl(db)
	tokenDao := dao2.NewUserTokenDaoImpl(db)
	service := service2.NewUserServiceImpl(dao, tokenDao)
	controller := NewUserLoginController(service)

	userRoute := route.Group("/user")
	{
		userRoute.POST("/login", controller.Login)
		userRoute.POST("/register", controller.Register)
		userRoute.POST("/refresh_token", controller.RefreshToken)
		userRoute.POST("/logout", middleware.JWTMiddleware(middleware.User), controller.Logout)
	}
}
