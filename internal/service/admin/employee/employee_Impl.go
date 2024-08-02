package employee

import (
	"context"
	employeeDao "sky-take-out-gin/internal/dao/admin/employee"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/employee"
)

type EmployeeServiceImpl struct {
	employeeDao.EmployeeDAO
}

func (service EmployeeServiceImpl) EditPassword(ctx context.Context, req employee.EditPasswordRequest) (*employee.EditPasswordResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) ChangeEmployeeStatus(ctx context.Context, request employee.ChangeEmployeeStatusRequest) (*employee.ChangeEmployeeStatusResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) GetEmployeePage(ctx context.Context, request employee.EmployeePageRequest) (*employee.EmployeePageResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) EmployeeLogin(ctx context.Context, request employee.EmployeeLoginRequest) (*employee.EmployeeLoginResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) AddEmployee(ctx context.Context, request employee.AddEmployeeRequest) (*employee.AddEmployeeResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) SearchEmployee(ctx context.Context, request employee.SearchEmployeeRequest) (*employee.SearchEmployeeResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) EditEmployee(ctx context.Context, request employee.EditEmployeeRequest) (*employee.EditEmployeeResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service EmployeeServiceImpl) EmployeeLogout(ctx context.Context, request employee.EmployeeLogoutRequest) (*employee.EmployeeLogoutResponse, *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func NewEmployeeService(employeeDao employeeDao.EmployeeDAO) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{employeeDao}
}
