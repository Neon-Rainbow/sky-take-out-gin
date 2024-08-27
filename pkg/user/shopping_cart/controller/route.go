package controller

import (
	"github.com/gin-gonic/gin"
	cache2 "sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
	dao2 "sky-take-out-gin/pkg/user/dish/dao"
	dao3 "sky-take-out-gin/pkg/user/set_meal/dao"
	"sky-take-out-gin/pkg/user/shopping_cart/dao"
	"sky-take-out-gin/pkg/user/shopping_cart/service"
)

func UserShoppingCartRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	userShoppingCartDao := dao.NewUserShoppingCartDatabaseImpl(db)
	userShoppingCartCache := dao.NewUserShoppingCartRedisImpl(db)
	cache := cache2.NewCache(db)
	userDishDao := dao2.NewDishDao(db)
	userSetMealDao := dao3.NewSetMealDaoImpl(db)

	userShoppingService := service.NewUserShoppingCartService(userShoppingCartCache, userShoppingCartDao, cache, userDishDao, userSetMealDao)

	userShoppingCartController := NewUserShoppingCartControllerImpl(userShoppingService)
	userShoppingCartRoute := route.Group("/shopping_cart").Use(middleware.JWTMiddleware(middleware.User))
	{
		userShoppingCartRoute.POST("/sub", userShoppingCartController.SubOneCommodity)
		userShoppingCartRoute.GET("/list", userShoppingCartController.GetShoppingCartList)
		userShoppingCartRoute.POST("/add", userShoppingCartController.AddToShoppingCart)
		userShoppingCartRoute.DELETE("/clean", userShoppingCartController.RemoveFromShoppingCart)
	}
}
