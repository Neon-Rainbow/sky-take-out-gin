package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/database"
	addressBookDao "sky-take-out-gin/pkg/user/address_book/dao"
	addressBookService "sky-take-out-gin/pkg/user/address_book/service"
)

func AddressBookRoute(route *gin.RouterGroup) {
	db := database.GetDatabaseManager()
	dao := addressBookDao.NewAddressBookDaoImpl(db)
	service := addressBookService.NewAddressBookServiceImpl(dao)
	controller := NewAddressBookControllerImpl(service)

	addressBookRoute := route.Group("/addressBook")
	{
		addressBookRoute.POST("", controller.AddAddressBook)
		addressBookRoute.GET("/list", controller.GetAddressBookList)
		addressBookRoute.GET("/default", controller.GetDefaultAddress)
		addressBookRoute.PUT("", controller.UpdateAddressBookByID)
		addressBookRoute.DELETE("", controller.DeleteAddressBookByID)
		addressBookRoute.GET("/:address_id", controller.GetAddressBookByID)
		addressBookRoute.PUT("/default", controller.SetDefaultAddress)
	}

}
