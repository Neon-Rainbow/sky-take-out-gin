package DTO

import (
	"sky-take-out-gin/model/sql"
)

type EditPasswordRequest struct {
	ID          uint   `json:"id" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

type EditPasswordResponse struct {
}

type ChangeEmployeeStatusRequest struct {
	Status int  `uri:"status" json:"status"`
	ID     uint `form:"id"`
}

type ChangeEmployeeStatusResponse struct {
}

type EmployeePageRequest struct {
	Name     string `form:"name" binding:"omitempty"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}

type EmployeePageResponse struct {
	Total   int              `json:"total"`
	Records []model.Employee `json:"records"`
}

type EmployeeLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeeLoginResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type AddEmployeeRequest struct {
	ID       uint   `json:"id" binding:"omitempty"`
	IdNumber string `json:"idNumber" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type AddEmployeeResponse struct {
}

type SearchEmployeeRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type SearchEmployeeResponse struct {
	model.Employee
}

type EditEmployeeRequest struct {
	ID       uint   `json:"id" binding:"omitempty"`
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
