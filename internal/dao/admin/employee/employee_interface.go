package employee

import model "sky-take-out-gin/model/sql"

// EmployeeDAO 员工数据访问接口
type EmployeeDAO interface {
	GetEmployeeByID(id int64) (*model.Employee, error)
	GetEmployees(page, pageSize int) ([]model.Employee, error)
	SearchEmployees(condition string, args ...interface{}) ([]model.Employee, error)
	UpdateEmployee(employee *model.Employee) error
	AddEmployee(employee *model.Employee) error
	DeleteEmployee(id int64) error
}
