package controller

import (
	"github.com/gin-gonic/gin"
	paramModel "sky-take-out-gin/pkg/admin/dish/DTO"
	dishService "sky-take-out-gin/pkg/admin/dish/service"
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
		controller.service.UpdateDish,
		c.ShouldBindBodyWithJSON,
	)
}

func (controller DishControllerImpl) DeleteDish(c *gin.Context) {
	req := paramModel.DeleteDishRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.service.DeleteDish,
		c.ShouldBindQuery,
	)
}

func (controller DishControllerImpl) AddDish(c *gin.Context) {
	req := paramModel.AddDishRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.service.AddDish,
		c.ShouldBindBodyWithJSON,
	)
}

func (controller DishControllerImpl) SearchDishByID(c *gin.Context) {
	req := paramModel.SearchDishByIDRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.service.SearchDishByID,
		c.ShouldBindUri,
	)
}

func (controller DishControllerImpl) SearchDishByCategory(c *gin.Context) {
	req := paramModel.SearchDishByCategoryRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.service.SearchDishByCategory,
		c.ShouldBindUri,
	)
}

func (controller DishControllerImpl) SearchDishByPage(c *gin.Context) {
	req := paramModel.SearchDishByPageRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.service.SearchDishByPage,
		c.ShouldBindQuery,
	)
}

func (controller DishControllerImpl) ChangeDishStatus(c *gin.Context) {
	req := paramModel.ChangeDishStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		controller.service.ChangeDishStatus,
		c.ShouldBindUri,
		c.ShouldBindQuery,
	)
}

func NewDishControllerImpl(service dishService.DishServiceInterface) DishControllerImpl {
	return DishControllerImpl{service: service}
}
