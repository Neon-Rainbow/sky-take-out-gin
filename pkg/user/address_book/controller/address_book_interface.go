package controller

import "github.com/gin-gonic/gin"

type AddressBookControllerInterface interface {
	AddAddressBook(c *gin.Context)
	GetAddressBookList(c *gin.Context)
	GetDefaultAddress(c *gin.Context)
	UpdateAddressBookByID(c *gin.Context)
	DeleteAddressBookByID(c *gin.Context)
	GetAddressBookByID(c *gin.Context)
	SetDefaultAddress(c *gin.Context)
}
