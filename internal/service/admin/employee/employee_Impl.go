package employee

import (
	"context"
	"sky-take-out-gin/code"
	employeeDao "sky-take-out-gin/internal/dao/admin/employee"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/employee"
	paramModel "sky-take-out-gin/model/sql"
	"sky-take-out-gin/utils"
	"sky-take-out-gin/utils/convert"
	"sky-take-out-gin/utils/jwt"
	"time"
)

type EmployeeServiceImpl struct {
	employeeDao.EmployeeDAO
}

func (service EmployeeServiceImpl) EditPassword(ctx context.Context, req employee.EditPasswordRequest) (*employee.EditPasswordResponse, *model.ApiError) {
	e, err := service.GetEmployeeByID(ctx, int64(req.ID))
	if err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeNotFound,
			Msg:  "Employee not found",
		}
	}

	e.Password = req.NewPassword
	if err := service.UpdateEmployee(ctx, e); err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeEditPasswordFailed,
			Msg:  "Failed to update password",
		}
	}
	return nil, nil
}

func (service EmployeeServiceImpl) ChangeEmployeeStatus(ctx context.Context, req employee.ChangeEmployeeStatusRequest) (*employee.ChangeEmployeeStatusResponse, *model.ApiError) {
	e, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeNotFound,
			Msg:  "Employee not found",
		}
	}

	e.Status = req.Status
	if err := service.UpdateEmployee(ctx, e); err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeChangeStatusFailed,
			Msg:  "Failed to change status",
		}
	}

	return nil, nil
}

func (service EmployeeServiceImpl) GetEmployeePage(ctx context.Context, req employee.EmployeePageRequest) (*employee.EmployeePageResponse, *model.ApiError) {
	employees, err := service.GetEmployees(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeGetPageFailed,
			Msg:  "Failed to get employee page",
		}
	}

	return &employee.EmployeePageResponse{
		Total:   int64(len(employees)),
		Records: employees,
	}, nil
}

func (service EmployeeServiceImpl) EmployeeLogin(ctx context.Context, req employee.EmployeeLoginRequest) (resp *employee.EmployeeLoginResponse, apiError *model.ApiError) {
	employees, err := service.SearchEmployees(ctx, req.Username, "username = ? AND password = ?", utils.EncryptPassword(req.Password))
	if err != nil || len(employees) == 0 {
		return nil, &model.ApiError{
			Code: code.EmployeeLoginFailed,
			Msg:  "Invalid username or password",
		}
	}
	employeeResult := employees[0]
	resp = &employee.EmployeeLoginResponse{}
	resp.ID = employeeResult.ID
	resp.Name = employeeResult.Name
	resp.Username = employeeResult.Username
	resp.Token, _, err = jwt.GenerateToken(employeeResult.Username, employeeResult.ID, true)

	return resp, nil
}

func (service EmployeeServiceImpl) AddEmployee(ctx context.Context, req employee.AddEmployeeRequest) (resp *employee.AddEmployeeResponse, apiError *model.ApiError) {
	var e paramModel.Employee
	err := convert.UpdateStructFields(&req, &e)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeAddFailed,
			Msg:  "Failed to add employee",
		}
	}

	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	e.Status = 1
	e.Password = utils.EncryptPassword("123456")
	err = service.EmployeeDAO.AddEmployee(ctx, &e)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeAddFailed,
			Msg:  "Failed to add employee",
		}
	}
	return nil, nil

}

func (service EmployeeServiceImpl) SearchEmployee(ctx context.Context, req employee.SearchEmployeeRequest) (resp *employee.SearchEmployeeResponse, apiError *model.ApiError) {
	employees, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.EmployeeSearchFailed,
			Msg:  "Failed to search employee",
		}
	}

	resp = &employee.SearchEmployeeResponse{
		Employee: *employees,
	}
	return resp, nil
}

func (service EmployeeServiceImpl) EditEmployee(ctx context.Context, req employee.EditEmployeeRequest) (resp *employee.EditEmployeeResponse, apiError *model.ApiError) {
	e, err := service.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		return nil, &model.ApiError{
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
		return nil, &model.ApiError{
			Code: code.EmployeeEditFailed,
			Msg:  "Failed to edit employee",
		}
	}

	return nil, nil
}

func (service EmployeeServiceImpl) EmployeeLogout(ctx context.Context, req employee.EmployeeLogoutRequest) (resp *employee.EmployeeLogoutResponse, apiError *model.ApiError) {
	// Implement logout logic here, if needed
	return nil, nil
}

func NewEmployeeService(employeeDao employeeDao.EmployeeDAO) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{employeeDao}
}
