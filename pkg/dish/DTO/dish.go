package DTO

import model "sky-take-out-gin/model/sql"

type UpdateDishRequest struct {
	model.Dish
}

type UpdateDishResponse struct {
}

type DeleteDishRequest struct {
	IDs []int64 `form:"ids" binding:"required"` // ids
}

type DeleteDishResponse struct {
}

type AddDishRequest struct {
	model.Dish
}

type AddDishResponse struct {
}

type SearchDishByIDRequest struct {
	ID int64 `uri:"id" binding:"required"` // 菜品id
}

type SearchDishByIDResponse struct {
	model.Dish
}

type SearchDishByCategoryRequest struct {
	CategoryID int64 `uri:"category_id" binding:"required"` // 分类id
}

type SearchDishByCategoryResponse struct {
	Records []model.Dish
}

type SearchDishByPageRequest struct {
	CategoryID int64  `form:"category_id"`                 // 分类id
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
	Status int   `uri:"status"` // 菜品状态，1表示上架，0表示下架
	ID     int64 `form:"id"`    // 菜品id
}

type ChangeDishStatusResponse struct {
}
