package service

import (
	"context"
	"sky-take-out-gin/internal/utils/convert"
	"sky-take-out-gin/internal/utils/encrypt"
	paramModel "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/JWT"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/employee/DTO"
	employeeDao "sky-take-out-gin/pkg/employee/dao"
)

type EmployeeServiceImpl struct {
	employeeDao.EmployeeDAOInterface
}

func (service EmployeeServiceImpl) EditPassword(ctx context.Context, req DTO.EditPasswordRequest) (*DTO.EditPasswordResponse, *error2.ApiError) {
	e, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeNotFound,
			Msg:  "Employee not found",
		}
	}

	e.Password = req.NewPassword
	if err := service.UpdateEmployee(ctx, e); err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeEditPasswordFailed,
			Msg:  "Failed to update password",
		}
	}
	return nil, nil
}

func (service EmployeeServiceImpl) ChangeEmployeeStatus(ctx context.Context, req DTO.ChangeEmployeeStatusRequest) (*DTO.ChangeEmployeeStatusResponse, *error2.ApiError) {
	e, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeNotFound,
			Msg:  "Employee not found",
		}
	}

	e.Status = req.Status
	if err := service.UpdateEmployee(ctx, e); err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeChangeStatusFailed,
			Msg:  "Failed to change status",
		}
	}

	return nil, nil
}

func (service EmployeeServiceImpl) GetEmployeePage(ctx context.Context, req DTO.EmployeePageRequest) (*DTO.EmployeePageResponse, *error2.ApiError) {
	employees, err := service.GetEmployees(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeGetPageFailed,
			Msg:  "Failed to get employee page",
		}
	}

	return &DTO.EmployeePageResponse{
		Total:   len(employees),
		Records: employees,
	}, nil
}

func (service EmployeeServiceImpl) EmployeeLogin(ctx context.Context, req DTO.EmployeeLoginRequest) (resp *DTO.EmployeeLoginResponse, apiError *error2.ApiError) {
	employees, err := service.SearchEmployees(ctx, "username = ?", req.Username)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeLoginFailed,
			Msg:  "Invalid username or password",
		}
	}
	if len(employees) == 0 {
		return nil, &error2.ApiError{
			Code: code.EmployeeLoginFailed,
			Msg:  "用户不存在",
		}
	}
	if encrypt.ComparePassword(req.Password, employees[0].Password) == false {
		return nil, &error2.ApiError{
			Code: code.EmployeeLoginFailed,
			Msg:  "密码错误",
		}
	}

	employeeResult := employees[0]
	resp = &DTO.EmployeeLoginResponse{}
	resp.ID = employeeResult.ID
	resp.Name = employeeResult.Name
	resp.Username = employeeResult.Username
	resp.Token, _, err = JWT.GenerateToken(employeeResult.Username, employeeResult.ID, true)

	return resp, nil
}

func (service EmployeeServiceImpl) AddEmployee(ctx context.Context, req DTO.AddEmployeeRequest) (resp *DTO.AddEmployeeResponse, apiError *error2.ApiError) {
	var e paramModel.Employee
	err := convert.UpdateStructFields(&req, &e)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeAddFailed,
			Msg:  "Failed to add employee",
		}
	}

	e.Status = 1
	e.Password = encrypt.EncryptPassword("123456")
	err = service.EmployeeDAOInterface.AddEmployee(ctx, &e)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeAddFailed,
			Msg:  "Failed to add employee",
		}
	}
	return nil, nil

}

func (service EmployeeServiceImpl) SearchEmployee(ctx context.Context, req DTO.SearchEmployeeRequest) (resp *DTO.SearchEmployeeResponse, apiError *error2.ApiError) {
	employees, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeSearchFailed,
			Msg:  "Failed to search employee",
		}
	}

	resp = &DTO.SearchEmployeeResponse{
		Employee: *employees,
	}
	return resp, nil
}

func (service EmployeeServiceImpl) EditEmployee(ctx context.Context, req DTO.EditEmployeeRequest) (resp *DTO.EditEmployeeResponse, apiError *error2.ApiError) {
	e, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeNotFound,
			Msg:  "Employee not found",
		}
	}

	err = convert.UpdateStructFields(&req, e)
	if err != nil {
		return nil, nil
	}

	err = service.UpdateEmployee(ctx, e)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.EmployeeEditFailed,
			Msg:  "Failed to edit employee",
		}
	}

	return nil, nil
}

func (service EmployeeServiceImpl) EmployeeLogout(ctx context.Context, req DTO.EmployeeLogoutRequest) (resp *DTO.EmployeeLogoutResponse, apiError *error2.ApiError) {
	// Implement logout logic here, if needed
	return nil, nil
}

func NewEmployeeService(employeeDao employeeDao.EmployeeDAOInterface) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{employeeDao}
}
