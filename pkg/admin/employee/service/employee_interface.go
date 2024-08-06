package service

import (
	"context"
	"sky-take-out-gin/pkg/admin/employee/DTO"
	apiErrorModel "sky-take-out-gin/pkg/common/error"
)

type EmployeeService interface {
	EditPassword(ctx context.Context, req DTO.EditPasswordRequest) (*DTO.EditPasswordResponse, *apiErrorModel.ApiError)
	ChangeEmployeeStatus(ctx context.Context, request DTO.ChangeEmployeeStatusRequest) (*DTO.ChangeEmployeeStatusResponse, *apiErrorModel.ApiError)
	GetEmployeePage(ctx context.Context, request DTO.EmployeePageRequest) (*DTO.EmployeePageResponse, *apiErrorModel.ApiError)
	EmployeeLogin(ctx context.Context, request DTO.EmployeeLoginRequest) (*DTO.EmployeeLoginResponse, *apiErrorModel.ApiError)
	AddEmployee(ctx context.Context, request DTO.AddEmployeeRequest) (*DTO.AddEmployeeResponse, *apiErrorModel.ApiError)
	SearchEmployee(ctx context.Context, request DTO.SearchEmployeeRequest) (*DTO.SearchEmployeeResponse, *apiErrorModel.ApiError)
	EditEmployee(ctx context.Context, request DTO.EditEmployeeRequest) (*DTO.EditEmployeeResponse, *apiErrorModel.ApiError)
	EmployeeLogout(ctx context.Context, request DTO.EmployeeLogoutRequest) (*DTO.EmployeeLogoutResponse, *apiErrorModel.ApiError)
}
