package controller

import (
	"github.com/gin-gonic/gin"
	employeeService "sky-take-out-gin/pkg/admin/employee/service"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
)

// EmployeeControllerImpl 员工控制器实现
type EmployeeControllerImpl struct {
	employeeService.EmployeeService
}

// NewEmployeeControllerImpl 创建员工控制器实例
func NewEmployeeControllerImpl(service employeeService.EmployeeService) *EmployeeControllerImpl {
	return &EmployeeControllerImpl{service}
}

func (controller *EmployeeControllerImpl) EditPassword(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.EditPassword)
}

func (controller *EmployeeControllerImpl) ChangeEmployeeStatus(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.ChangeEmployeeStatus, c.ShouldBindQuery, c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EmployeePage(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.GetEmployeePage)
}

func (controller *EmployeeControllerImpl) EmployeeLogin(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.EmployeeLogin)
}

func (controller *EmployeeControllerImpl) AddEmployee(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.AddEmployee)
}

func (controller *EmployeeControllerImpl) SearchEmployee(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.SearchEmployee, c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EditEmployee(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.EditEmployee)
}

func (controller *EmployeeControllerImpl) EmployeeLogout(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.EmployeeService.EmployeeLogout)
}
