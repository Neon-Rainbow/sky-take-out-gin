package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	paramModel "sky-take-out-gin/pkg/admin/dish/DTO"
	dishService "sky-take-out-gin/pkg/admin/dish/service"
	error2 "sky-take-out-gin/pkg/common/error"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
)

type DishControllerImpl struct {
	service dishService.DishServiceInterface
}

func (controller DishControllerImpl) UpdateDish(c *gin.Context) {
	req := paramModel.UpdateDishRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.UpdateDish(ctx, req.(*paramModel.UpdateDishRequest))
		},
		c.ShouldBindBodyWithJSON,
	)
}

func (controller DishControllerImpl) DeleteDish(c *gin.Context) {
	req := paramModel.DeleteDishRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.DeleteDish(ctx, req.(*paramModel.DeleteDishRequest))
		},
		c.ShouldBindQuery,
	)
}

func (controller DishControllerImpl) AddDish(c *gin.Context) {
	req := paramModel.AddDishRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.AddDish(ctx, req.(*paramModel.AddDishRequest))
		},
		c.ShouldBindBodyWithJSON,
	)
}

func (controller DishControllerImpl) SearchDishByID(c *gin.Context) {
	req := paramModel.SearchDishByIDRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.SearchDishByID(ctx, req.(*paramModel.SearchDishByIDRequest))
		},
		c.ShouldBindUri,
	)
}

func (controller DishControllerImpl) SearchDishByCategory(c *gin.Context) {
	req := paramModel.SearchDishByCategoryRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.SearchDishByCategory(ctx, req.(*paramModel.SearchDishByCategoryRequest))
		},
		c.ShouldBindUri,
	)
}

func (controller DishControllerImpl) SearchDishByPage(c *gin.Context) {
	req := paramModel.SearchDishByPageRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.SearchDishByPage(ctx, req.(*paramModel.SearchDishByPageRequest))
		},
		c.ShouldBindQuery,
	)
}

func (controller DishControllerImpl) ChangeDishStatus(c *gin.Context) {
	req := paramModel.ChangeDishStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError) {
			return controller.service.ChangeDishStatus(ctx, req.(*paramModel.ChangeDishStatusRequest))
		},
		c.ShouldBindUri,
		c.ShouldBindQuery,
	)
}

func NewDishControllerImpl(service dishService.DishServiceInterface) DishControllerImpl {
	return DishControllerImpl{service: service}
}