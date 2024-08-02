package employee

import (
	"gorm.io/gorm"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/utils/MySQL"
)

// EmployeeDAOImpl 员工数据访问接口实现
type EmployeeDAOImpl struct {
	db *gorm.DB
}

// NewEmployeeDAOImpl 实例化EmployeeDAOImpl
func NewEmployeeDAOImpl() *EmployeeDAOImpl {
	return &EmployeeDAOImpl{MySQL.GetDB()}
}

func (dao EmployeeDAOImpl) GetEmployeeByID(id int64) (*model.Employee, error) {
	var employee model.Employee
	if err := dao.db.First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (dao EmployeeDAOImpl) GetEmployees(page, pageSize int) ([]model.Employee, error) {
	var employees []model.Employee
	offset := (page - 1) * pageSize
	if err := dao.db.Offset(offset).Limit(pageSize).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// SearchEmployees 搜索员工
// @Param condition string 搜索条件 样例: "name = ?"
// @Param args ...interface{} 搜索条件参数 样例: "张三"
// @Return []model.Employee 员工列表
// @Return error 错误信息
func (dao EmployeeDAOImpl) SearchEmployees(condition string, args ...interface{}) ([]model.Employee, error) {
	var employees []model.Employee
	if err := dao.db.Where(condition, args...).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (dao EmployeeDAOImpl) UpdateEmployee(employee *model.Employee) error {
	return dao.db.Save(employee).Error
}

func (dao EmployeeDAOImpl) AddEmployee(employee *model.Employee) error {
	return dao.db.Create(employee).Error
}

func (dao EmployeeDAOImpl) DeleteEmployee(id int64) error {
	return dao.db.Delete(&model.Employee{}, id).Error
}
