package controller

import (
	"github.com/gin-gonic/gin"
	setmealService "sky-take-out-gin/pkg/admin/setmeal/service"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
)

type SetmealControllerImpl struct {
	service setmealService.SetmealServiceInterface
}

func NewSetmealControllerImpl(service setmealService.SetmealServiceInterface) SetmealControllerImpl {
	return SetmealControllerImpl{service}
}

func (controller SetmealControllerImpl) UpdateSetmeal(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.UpdateSetmeal, c.ShouldBindBodyWithJSON)
}

func (controller SetmealControllerImpl) GetSetmealPage(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.GetSetmealPage, c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) ChangeSetmealStatus(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.ChangeSetmealStatus, c.ShouldBindUri, c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) DeleteSetmeals(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.DeleteSetmeals, c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) CreateSetmeals(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.CreateSetmeals, c.ShouldBindBodyWithJSON)
}

func (controller SetmealControllerImpl) GetSetmealsByID(c *gin.Context) {

	HandleRequest.HandleRequest(c, controller.service.GetSetmealsByID, c.ShouldBindUri)
}
