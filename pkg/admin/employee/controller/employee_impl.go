package controller

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/admin/employee/DTO"
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
	req := DTO.EditPasswordRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.EditPassword)
}

func (controller *EmployeeControllerImpl) ChangeEmployeeStatus(c *gin.Context) {
	req := DTO.ChangeEmployeeStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.EmployeeService.ChangeEmployeeStatus,
		c.ShouldBindQuery,
		c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EmployeePage(c *gin.Context) {
	req := DTO.EmployeePageRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.GetEmployeePage)
}

func (controller *EmployeeControllerImpl) EmployeeLogin(c *gin.Context) {
	req := DTO.EmployeeLoginRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.EmployeeLogin)
}

func (controller *EmployeeControllerImpl) AddEmployee(c *gin.Context) {
	req := DTO.AddEmployeeRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.AddEmployee)
}

func (controller *EmployeeControllerImpl) SearchEmployee(c *gin.Context) {
	req := DTO.SearchEmployeeRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.SearchEmployee, c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EditEmployee(c *gin.Context) {
	req := DTO.EditEmployeeRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.EditEmployee)
}

func (controller *EmployeeControllerImpl) EmployeeLogout(c *gin.Context) {
	req := DTO.EmployeeLogoutRequest{}
	HandleRequest.HandleRequest(c, &req, controller.EmployeeService.EmployeeLogout)
}
