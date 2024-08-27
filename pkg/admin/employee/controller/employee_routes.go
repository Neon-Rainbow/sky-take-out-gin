package controller

import (
	"github.com/gin-gonic/gin"
	employeeDao "sky-take-out-gin/pkg/admin/employee/dao"
	"sky-take-out-gin/pkg/admin/employee/service"
	"sky-take-out-gin/pkg/common/database"
	"sky-take-out-gin/pkg/common/middleware"
)

func EmployeeRoutes(route *gin.RouterGroup) {
	// 实例化EmployeeServiceImpl
	db := database.GetDatabaseManager()
	dao := employeeDao.NewEmployeeDAOImpl(db)
	service := service.NewEmployeeService(dao)
	controller := NewEmployeeControllerImpl(service)

	employeeRoute := route.Group("/employee")
	{
		employeeRoute.PUT("/password", middleware.JWTMiddleware(middleware.Admin), controller.EditPassword)
		employeeRoute.POST("/status/:status", middleware.JWTMiddleware(middleware.Admin), controller.ChangeEmployeeStatus)
		employeeRoute.GET("/page", middleware.JWTMiddleware(middleware.Admin), controller.EmployeePage)
		employeeRoute.POST("/login", controller.EmployeeLogin)
		employeeRoute.POST("/", middleware.JWTMiddleware(middleware.Admin), controller.AddEmployee)
		employeeRoute.GET("/:id", middleware.JWTMiddleware(middleware.Admin), controller.SearchEmployee)
		employeeRoute.PUT("/", middleware.JWTMiddleware(middleware.Admin), controller.EditEmployee)
		employeeRoute.POST("/logout", middleware.JWTMiddleware(middleware.Admin), controller.EmployeeLogout)
	}
}
