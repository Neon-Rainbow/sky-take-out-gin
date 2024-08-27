package controller

import (
	"github.com/gin-gonic/gin"
	paramModel "sky-take-out-gin/pkg/admin/setmeal/DTO"
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
	var req paramModel.UpdateSetmealRequest
	HandleRequest.HandleRequest(c, &req, controller.service.UpdateSetmeal, c.ShouldBindBodyWithJSON)
}

func (controller SetmealControllerImpl) GetSetmealPage(c *gin.Context) {
	var req paramModel.GetSetmealsPageRequest
	HandleRequest.HandleRequest(c, &req, controller.service.GetSetmealPage, c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) ChangeSetmealStatus(c *gin.Context) {
	var req paramModel.UpdateSetmealStatusRequest
	HandleRequest.HandleRequest(c, &req, controller.service.ChangeSetmealStatus, c.ShouldBindUri, c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) DeleteSetmeals(c *gin.Context) {
	var req paramModel.DeleteSetmealsRequest
	HandleRequest.HandleRequest(c, &req, controller.service.DeleteSetmeals, c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) CreateSetmeals(c *gin.Context) {
	var req paramModel.AddSetmealRequest
	HandleRequest.HandleRequest(c, &req, controller.service.CreateSetmeals, c.ShouldBindBodyWithJSON)
}

func (controller SetmealControllerImpl) GetSetmealsByID(c *gin.Context) {
	var req paramModel.GetSetmealByIDRequest
	HandleRequest.HandleRequest(c, &req, controller.service.GetSetmealsByID, c.ShouldBindUri)
}
