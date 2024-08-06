package DTO

import (
	model "sky-take-out-gin/model/sql"
)

type DishFlavorDTO struct {
	DishID uint     `json:"dishId"`
	ID     uint     `json:"id"`
	Name   string   `json:"name" binding:"required"`
	Value  []string `json:"value" binding:"required"`
}

type UpdateDishRequest struct {
	ID          uint            `json:"id" binding:"required"` // 菜品id
	Name        string          `json:"name"`                  // 菜品名称
	CategoryID  uint            `json:"category_id"`           // 分类id
	Price       float64         `json:"price"`                 // 价格
	Image       string          `json:"image"`                 // 图片
	Description string          `json:"description"`           // 描述
	Status      int             `json:"status"`                // 状态
	Flavors     []DishFlavorDTO `json:"flavors"`
}

type UpdateDishResponse struct {
}

type DeleteDishRequest struct {
	IDs []uint `form:"ids" binding:"required"` // ids
}

type DeleteDishResponse struct {
}

type AddDishRequest struct {
	ID          uint            `json:"id"`          // 菜品id
	Name        string          `json:"name"`        // 菜品名称
	CategoryID  uint            `json:"category_id"` // 分类id
	Price       float64         `json:"price"`       // 价格
	Image       string          `json:"image"`       // 图片
	Description string          `json:"description"` // 描述
	Status      int             `json:"status"`      // 状态
	Flavors     []DishFlavorDTO `json:"flavors"`
}

type AddDishResponse struct {
}

type SearchDishByIDRequest struct {
	ID uint `uri:"id" binding:"required"` // 菜品id
}

type SearchDishByIDResponse struct {
	model.Dish
}

type SearchDishByCategoryRequest struct {
	CategoryID uint `uri:"category_id" binding:"required"` // 分类id
}

type SearchDishByCategoryResponse struct {
	Records []model.Dish
}

type SearchDishByPageRequest struct {
	CategoryID uint   `form:"category_id"`                 // 分类id
	Name       string `form:"name"`                        // 菜品名称
	Status     int    `form:"status"`                      // 菜品状态
	Page       int    `form:"page" binding:"required"`     // 页码
	PageSize   int    `form:"pageSize" binding:"required"` // 每页记录数
}

type SearchDishByPageResponse struct {
	Total   int          `json:"total"`
	Records []model.Dish `json:"records"`
}

type ChangeDishStatusRequest struct {
	Status int  `uri:"status"` // 菜品状态，1表示上架，0表示下架
	ID     uint `form:"id"`    // 菜品id
}

type ChangeDishStatusResponse struct {
}
