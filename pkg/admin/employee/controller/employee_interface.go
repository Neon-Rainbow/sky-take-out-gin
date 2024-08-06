package controller

import "github.com/gin-gonic/gin"

type EmployeeController interface {
	EditPassword(c *gin.Context)
	ChangeEmployeeStatus(c *gin.Context)
	EmployeePage(c *gin.Context)
	EmployeeLogin(c *gin.Context)
	AddEmployee(c *gin.Context)
	SearchEmployee(c *gin.Context)
	EditEmployee(c *gin.Context)
	EmployeeLogout(c *gin.Context)
}
