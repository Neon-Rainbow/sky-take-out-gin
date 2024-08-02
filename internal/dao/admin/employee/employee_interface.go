package employee

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

// EmployeeDAO 员工数据访问接口
type EmployeeDAO interface {
	GetEmployeeByID(ctx context.Context, id int64) (*model.Employee, error)
	GetEmployees(ctx context.Context, page, pageSize int) ([]model.Employee, error)
	SearchEmployees(ctx context.Context, condition string, args ...interface{}) ([]model.Employee, error)
	UpdateEmployee(ctx context.Context, employee *model.Employee) error
	AddEmployee(ctx context.Context, employee *model.Employee) error
	DeleteEmployee(ctx context.Context, id int64) error
}
