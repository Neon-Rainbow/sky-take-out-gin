package employee

import (
	"context"
	"sky-take-out-gin/model/param/admin/employee"
)
import apiErrorModel "sky-take-out-gin/model"

type EmployeeService interface {
	EditPassword(ctx context.Context, req employee.EditPasswordRequest) (*employee.EditPasswordResponse, *apiErrorModel.ApiError)
	ChangeEmployeeStatus(ctx context.Context, request employee.ChangeEmployeeStatusRequest) (*employee.ChangeEmployeeStatusResponse, *apiErrorModel.ApiError)
	GetEmployeePage(ctx context.Context, request employee.EmployeePageRequest) (*employee.EmployeePageResponse, *apiErrorModel.ApiError)
	EmployeeLogin(ctx context.Context, request employee.EmployeeLoginRequest) (*employee.EmployeeLoginResponse, *apiErrorModel.ApiError)
	AddEmployee(ctx context.Context, request employee.AddEmployeeRequest) (*employee.AddEmployeeResponse, *apiErrorModel.ApiError)
	SearchEmployee(ctx context.Context, request employee.SearchEmployeeRequest) (*employee.SearchEmployeeResponse, *apiErrorModel.ApiError)
	EditEmployee(ctx context.Context, request employee.EditEmployeeRequest) (*employee.EditEmployeeResponse, *apiErrorModel.ApiError)
	EmployeeLogout(ctx context.Context, request employee.EmployeeLogoutRequest) (*employee.EmployeeLogoutResponse, *apiErrorModel.ApiError)
}
