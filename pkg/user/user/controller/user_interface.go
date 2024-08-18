package controller

import "github.com/gin-gonic/gin"

type UserLoginControllerInterface interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Register(c *gin.Context)
	RefreshToken(c *gin.Context)
}
