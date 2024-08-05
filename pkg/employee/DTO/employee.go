package employee

import model "sky-take-out-gin/model/sql"

type EditPasswordRequest struct {
	ID          int    `json:"id" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

type EditPasswordResponse struct {
}

type ChangeEmployeeStatusRequest struct {
	Status int   `uri:"status" json:"status"`
	ID     int64 `form:"id"`
}

type ChangeEmployeeStatusResponse struct {
}

type EmployeePageRequest struct {
	Name     string `form:"name" binding:"omitempty"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}

type EmployeePageResponse struct {
	Total   int64            `json:"total"`
	Records []model.Employee `json:"records"`
}

type EmployeeLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeeLoginResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type AddEmployeeRequest struct {
	ID       int64  `json:"id" binding:"omitempty"`
	IdNumber string `json:"idNumber" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type AddEmployeeResponse struct {
}

type SearchEmployeeRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type SearchEmployeeResponse struct {
	model.Employee
}

type EditEmployeeRequest struct {
	ID       int64  `json:"id" binding:"omitempty"`
	IdNumber string `json:"idNumber" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type EditEmployeeResponse struct {
}

type EmployeeLogoutRequest struct {
}

type EmployeeLogoutResponse struct {
}
