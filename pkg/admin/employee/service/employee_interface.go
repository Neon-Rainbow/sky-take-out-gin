package service

import (
	"context"
	"sky-take-out-gin/pkg/admin/employee/DTO"
	apiErrorModel "sky-take-out-gin/pkg/common/api_error"
)

type EmployeeService interface {
	EditPassword(ctx context.Context, req *DTO.EditPasswordRequest) (*DTO.EditPasswordResponse, *apiErrorModel.ApiError)
	ChangeEmployeeStatus(ctx context.Context, req *DTO.ChangeEmployeeStatusRequest) (*DTO.ChangeEmployeeStatusResponse, *apiErrorModel.ApiError)
	GetEmployeePage(ctx context.Context, req *DTO.EmployeePageRequest) (*DTO.EmployeePageResponse, *apiErrorModel.ApiError)
	EmployeeLogin(ctx context.Context, req *DTO.EmployeeLoginRequest) (*DTO.EmployeeLoginResponse, *apiErrorModel.ApiError)
	AddEmployee(ctx context.Context, req *DTO.AddEmployeeRequest) (*DTO.AddEmployeeResponse, *apiErrorModel.ApiError)
	SearchEmployee(ctx context.Context, req *DTO.SearchEmployeeRequest) (*DTO.SearchEmployeeResponse, *apiErrorModel.ApiError)
	EditEmployee(ctx context.Context, req *DTO.EditEmployeeRequest) (*DTO.EditEmployeeResponse, *apiErrorModel.ApiError)
	EmployeeLogout(ctx context.Context, req *DTO.EmployeeLogoutRequest) (*DTO.EmployeeLogoutResponse, *apiErrorModel.ApiError)
}
