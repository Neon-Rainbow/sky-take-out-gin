package DTO

import (
	model "sky-take-out-gin/model/sql"
)

// AdminUpdateCategoryRequest 更新分类请求
type AdminUpdateCategoryRequest struct {
	// 分类ID
	ID uint `json:"id" binding:"required"`

	// 分类名称
	Name string `json:"name" binding:"required"`

	// 分类类型
	Type int `json:"type" binding:"required"`

	// 分类排序
	Sort int `json:"sort" binding:"required"`
}

// AdminUpdateCategoryResponse 更新分类响应
type AdminUpdateCategoryResponse struct {
}

// AdminCategoryPageRequest 分类分页查询请求
type AdminCategoryPageRequest struct {
	// 分类名称
	Name string `form:"name" binding:"omitempty"`

	// 页码
	Page int `form:"page" binding:"required"`

	// 每页记录数
	PageSize int `form:"pageSize" binding:"required"`

	// 分类类型：1为菜品分类，2为套餐分类
	Type int `form:"type" binding:"omitempty"`
}

// CategoryRecord 分类记录
type CategoryRecord struct {
	ID         uint   `json:"id"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	CreateUser int    `json:"createUser"`
	UpdateUser int    `json:"updateUser"`
}

// AdminCategoryPageResponse 分类分页查询响应
type AdminCategoryPageResponse struct {
	Total   int64            `json:"total"`
	Records []CategoryRecord `json:"records"`
}

// AdminChangeCategoryStatusRequest 启用、禁用分类请求
type AdminChangeCategoryStatusRequest struct {
	// 分类状态 1:启用 2:禁用
	Status int `uri:"status"`
	// 分类ID
	ID uint `form:"id"`
}

// AdminChangeCategoryStatusResponse 启用、禁用分类响应
type AdminChangeCategoryStatusResponse struct {
}

// AdminCreateCategoryRequest 新增分类请求
type AdminCreateCategoryRequest struct {
	// 分类名称
	Name string `json:"name" binding:"required"`

	// 分类类型
	Type int `json:"type" binding:"required"`

	// 分类排序
	Sort int `json:"sort" binding:"required"`
}

// AdminCreateCategoryResponse 新增分类响应
type AdminCreateCategoryResponse struct {
}

// AdminDeleteCategoryRequest 删除分类请求
type AdminDeleteCategoryRequest struct {
	// 分类ID
	ID uint `form:"id" binding:"required"`
}

// AdminDeleteCategoryResponse 删除分类响应
type AdminDeleteCategoryResponse struct {
}

// AdminGetCategoryListByTypeRequest 根据类型查询分类请求
type AdminGetCategoryListByTypeRequest struct {
	// 分类类型：1为菜品分类，2为套餐分类
	Type int `form:"type" binding:"omitempty"`
}

// AdminGetCategoryListByTypeResponse 根据类型查询分类响应
type AdminGetCategoryListByTypeResponse struct {
	CategoryList []model.Category `json:"categoryList"`
}
