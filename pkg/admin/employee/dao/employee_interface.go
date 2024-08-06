package dao

import (
	"context"
	"sky-take-out-gin/model/sql"
)

// EmployeeDAOInterface 员工数据访问接口
type EmployeeDAOInterface interface {
	GetEmployeeByID(ctx context.Context, id uint) (*model.Employee, error)
	GetEmployees(ctx context.Context, page, pageSize int) ([]model.Employee, error)
	SearchEmployees(ctx context.Context, condition string, args ...interface{}) ([]model.Employee, error)
	UpdateEmployee(ctx context.Context, employee *model.Employee) error
	AddEmployee(ctx context.Context, employee *model.Employee) error
	DeleteEmployee(ctx context.Context, id uint) error
}
