package employee

import (
	"context"
	"github.com/gin-gonic/gin"
	HandleRequest "sky-take-out-gin/internal/controller"
	employeeService "sky-take-out-gin/internal/service/admin/employee"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/employee"
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
	req := employee.EditPasswordRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.EditPassword(ctx, *req.(*employee.EditPasswordRequest))
		})

}

func (controller *EmployeeControllerImpl) ChangeEmployeeStatus(c *gin.Context) {
	req := employee.ChangeEmployeeStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.ChangeEmployeeStatus(ctx, *req.(*employee.ChangeEmployeeStatusRequest))
		},
		c.ShouldBindQuery,
		c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EmployeePage(c *gin.Context) {
	req := employee.EmployeePageRequest{}
	HandleRequest.HandleRequest(c, &req, func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
		return controller.EmployeeService.GetEmployeePage(ctx, *req.(*employee.EmployeePageRequest))
	})
}

func (controller *EmployeeControllerImpl) EmployeeLogin(c *gin.Context) {
	req := employee.EmployeeLoginRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.EmployeeLogin(ctx, *req.(*employee.EmployeeLoginRequest))
		})
}

func (controller *EmployeeControllerImpl) AddEmployee(c *gin.Context) {
	req := employee.AddEmployeeRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.AddEmployee(ctx, *req.(*employee.AddEmployeeRequest))
		})
}

func (controller *EmployeeControllerImpl) SearchEmployee(c *gin.Context) {
	req := employee.SearchEmployeeRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.SearchEmployee(ctx, *req.(*employee.SearchEmployeeRequest))
		},
		c.ShouldBindUri)
}

func (controller *EmployeeControllerImpl) EditEmployee(c *gin.Context) {
	req := employee.EditEmployeeRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.EditEmployee(ctx, *req.(*employee.EditEmployeeRequest))
		})
}

func (controller *EmployeeControllerImpl) EmployeeLogout(c *gin.Context) {
	req := employee.EmployeeLogoutRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.EmployeeService.EmployeeLogout(ctx, *req.(*employee.EmployeeLogoutRequest))
		})
}
