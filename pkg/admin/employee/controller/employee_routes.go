package controller

import (
	"github.com/gin-gonic/gin"
	employeeDao "sky-take-out-gin/pkg/admin/employee/dao"
	"sky-take-out-gin/pkg/admin/employee/service"
	"sky-take-out-gin/pkg/common/database"
)

func EmployeeRoutes(route *gin.RouterGroup) {
	// 实例化EmployeeServiceImpl
	db := database.GetDatabaseManager()
	dao := employeeDao.NewEmployeeDAOImpl(db)
	service := service.NewEmployeeService(dao)
	controller := NewEmployeeControllerImpl(service)

	employeeRoute := route.Group("/employee")
	{
		employeeRoute.PUT("/password", controller.EditPassword)
		employeeRoute.POST("/status/:status", controller.ChangeEmployeeStatus)
		employeeRoute.GET("/page", controller.EmployeePage)
		employeeRoute.POST("/login", controller.EmployeeLogin)
		employeeRoute.POST("/", controller.AddEmployee)
		employeeRoute.GET("/:id", controller.SearchEmployee)
		employeeRoute.PUT("/", controller.EditEmployee)
		employeeRoute.POST("/logout", controller.EmployeeLogout)
	}
}
