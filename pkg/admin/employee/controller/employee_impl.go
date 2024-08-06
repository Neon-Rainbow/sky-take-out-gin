package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/admin/employee/DTO"
	employeeService "sky-take-out-gin/pkg/admin/employee/service"
	error2 "sky-take-out-gin/pkg/common/error"
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
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.EditPassword(ctx, *req.(*DTO.EditPasswordRequest))
		})

}

func (controller *EmployeeControllerImpl) ChangeEmployeeStatus(c *gin.Context) {
	req := DTO.ChangeEmployeeStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.ChangeEmployeeStatus(ctx, *req.(*DTO.ChangeEmployeeStatusRequest))
		},
		c.ShouldBindQuery,
		c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EmployeePage(c *gin.Context) {
	req := DTO.EmployeePageRequest{}
	HandleRequest.HandleRequest(c, &req, func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
		return controller.EmployeeService.GetEmployeePage(ctx, *req.(*DTO.EmployeePageRequest))
	})
}

func (controller *EmployeeControllerImpl) EmployeeLogin(c *gin.Context) {
	req := DTO.EmployeeLoginRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.EmployeeLogin(ctx, *req.(*DTO.EmployeeLoginRequest))
		})
}

func (controller *EmployeeControllerImpl) AddEmployee(c *gin.Context) {
	req := DTO.AddEmployeeRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.AddEmployee(ctx, *req.(*DTO.AddEmployeeRequest))
		})
}

func (controller *EmployeeControllerImpl) SearchEmployee(c *gin.Context) {
	req := DTO.SearchEmployeeRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.SearchEmployee(ctx, *req.(*DTO.SearchEmployeeRequest))
		},
		c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EditEmployee(c *gin.Context) {
	req := DTO.EditEmployeeRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.EditEmployee(ctx, *req.(*DTO.EditEmployeeRequest))
		})
}

func (controller *EmployeeControllerImpl) EmployeeLogout(c *gin.Context) {
	req := DTO.EmployeeLogoutRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return controller.EmployeeService.EmployeeLogout(ctx, *req.(*DTO.EmployeeLogoutRequest))
		})
}
