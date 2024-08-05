package DTO

import (
	model "sky-take-out-gin/model/sql"
	"time"
)

type UpdateSetmealRequest struct {
	ID            int64               `json:"id" binding:"required"`          // 套餐id
	CategoryID    int64               `json:"category_id" binding:"required"` // 分类id
	Name          string              `json:"name" binding:"required"`        // 套餐名称
	Price         float64             `json:"price" binding:"required"`       // 套餐价格
	Status        int                 `json:"status"`                         // 套餐起售状态
	Description   string              `json:"description"`                    // 套餐描述
	Image         string              `json:"image"`                          // 套餐图片路径
	SetmealDishes []model.SetmealDish `json:"setmeal_dishes"`                 // 套餐和菜品关联关系
}

type UpdateSetmealResponse struct {
}

type GetSetmealsPageRequest struct {
	CategoryID int64  `form:"categoryId"`                  // 分类id
	Name       string `form:"name"`                        // 套餐名称
	Status     int    `form:"status"`                      // 套餐起售状态
	Page       int    `form:"page" binding:"required"`     // 页码
	PageSize   int    `form:"pageSize" binding:"required"` // 每页记录数
}

type GetSetmealsPageResponse struct {
	Total   int             `json:"total"`
	Records []model.Setmeal `json:"records"`
}

type UpdateSetmealStatusRequest struct {
	Status int   `uri:"status"` // 套餐状态，1表示起售，0表示停售
	ID     int64 `form:"id"`    // 套餐id
}

type UpdateSetmealStatusResponse struct {
}

type DeleteSetmealsRequest struct {
	IDs []int64 `form:"ids" binding:"required"` // ids
}

type DeleteSetmealsResponse struct {
}

type AddSetmealRequest struct {
	CategoryID    int64           `json:"category_id" binding:"required"`    // 分类id
	Name          string          `json:"name" binding:"required"`           // 套餐名称
	Price         float64         `json:"price" binding:"required"`          // 套餐价格
	Status        int             `json:"status" binding:"required"`         // 套餐状态：1为起售 0为停售
	Description   string          `json:"description"`                       // 套餐描述
	Image         string          `json:"image"`                             // 套餐图片
	SetmealDishes []model.Setmeal `json:"setmeal_dishes" binding:"required"` // 套餐和菜品关联关系
}

type AddSetmealResponse struct {
}

type GetSetmealByIDRequest struct {
	ID int64 `uri:"id" binding:"required"` // 套餐id
}

type GetSetmealByIDResponse struct {
	ID            int64               `json:"id"`
	CategoryID    int64               `json:"category_id"`
	Name          string              `json:"name"`
	Price         float64             `json:"price"`
	Status        int                 `json:"status"`
	Description   string              `json:"description"`
	Image         string              `json:"image"`
	UpdateTime    time.Time           `json:"update_time"`
	CategoryName  string              `json:"category_name"`
	SetmealDishes []model.SetmealDish `json:"setmeal_dishes"`
}
