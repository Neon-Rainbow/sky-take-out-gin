package controller

import (
	"github.com/gin-gonic/gin"
	employeeDao "sky-take-out-gin/pkg/employee/dao"
	"sky-take-out-gin/pkg/employee/service"
)

func EmployeeRoutes(route *gin.RouterGroup) {
	// 实例化EmployeeServiceImpl
	dao := employeeDao.NewEmployeeDAOImpl()
	service := service.NewEmployeeService(dao)
	controller := NewEmployeeControllerImpl(service)

	employee_route := route.Group("/employee")
	{
		employee_route.PUT("/password", controller.EditPassword)
		employee_route.POST("/status/:status", controller.ChangeEmployeeStatus)
		employee_route.GET("/page", controller.EmployeePage)
		employee_route.POST("/login", controller.EmployeeLogin)
		employee_route.POST("/", controller.AddEmployee)
		employee_route.GET("/:id", controller.SearchEmployee)
		employee_route.PUT("/", controller.EditEmployee)
		employee_route.POST("/logout", controller.EmployeeLogout)
	}
}
