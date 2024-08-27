package controller

import (
	"github.com/gin-gonic/gin"
	dishService "sky-take-out-gin/pkg/admin/dish/service"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
)

type DishControllerImpl struct {
	service dishService.DishServiceInterface
}

func (controller DishControllerImpl) UpdateDish(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.UpdateDish, c.ShouldBindBodyWithJSON)
}

func (controller DishControllerImpl) DeleteDish(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.DeleteDish, c.ShouldBindQuery)
}

func (controller DishControllerImpl) AddDish(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.AddDish, c.ShouldBindBodyWithJSON)
}

func (controller DishControllerImpl) SearchDishByID(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.SearchDishByID, c.ShouldBindUri)
}

func (controller DishControllerImpl) SearchDishByCategory(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.SearchDishByCategory, c.ShouldBindUri)
}

func (controller DishControllerImpl) SearchDishByPage(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.SearchDishByPage, c.ShouldBindQuery)
}

func (controller DishControllerImpl) ChangeDishStatus(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.ChangeDishStatus, c.ShouldBindUri, c.ShouldBindQuery)
}

func NewDishControllerImpl(service dishService.DishServiceInterface) DishControllerImpl {
	return DishControllerImpl{service: service}
}
