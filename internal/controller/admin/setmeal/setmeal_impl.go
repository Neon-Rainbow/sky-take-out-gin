package setmeal

import (
	"context"
	"github.com/gin-gonic/gin"
	HandleRequest "sky-take-out-gin/internal/controller"
	setmealService "sky-take-out-gin/internal/service/admin/setmeal"
	"sky-take-out-gin/model"
	paramModel "sky-take-out-gin/model/param/admin/setmeal"
)

type SetmealControllerImpl struct {
	service setmealService.SetmealServiceInterface
}

func NewSetmealControllerImpl(service setmealService.SetmealServiceInterface) SetmealControllerImpl {
	return SetmealControllerImpl{service}
}

func (controller SetmealControllerImpl) UpdateSetmeal(c *gin.Context) {
	var req paramModel.UpdateSetmealRequest
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.service.UpdateSetmeal(ctx, req.(*paramModel.UpdateSetmealRequest))
		},
		c.ShouldBindBodyWithJSON)
}

func (controller SetmealControllerImpl) GetSetmealPage(c *gin.Context) {
	var req paramModel.GetSetmealsPageRequest
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.service.GetSetmealPage(ctx, req.(*paramModel.GetSetmealsPageRequest))
		},
		c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) ChangeSetmealStatus(c *gin.Context) {
	var req paramModel.UpdateSetmealStatusRequest
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.service.ChangeSetmealStatus(ctx, req.(*paramModel.UpdateSetmealStatusRequest))
		},
		c.ShouldBindUri,
		c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) DeleteSetmeals(c *gin.Context) {
	var req paramModel.DeleteSetmealsRequest
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.service.DeleteSetmeals(ctx, req.(*paramModel.DeleteSetmealsRequest))
		},
		c.ShouldBindQuery)
}

func (controller SetmealControllerImpl) CreateSetmeals(c *gin.Context) {
	var req paramModel.AddSetmealRequest
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.service.CreateSetmeals(ctx, req.(*paramModel.AddSetmealRequest))
		},
		c.ShouldBindBodyWithJSON)
}

func (controller SetmealControllerImpl) GetSetmealsByID(c *gin.Context) {
	var req paramModel.GetSetmealByIDRequest
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return controller.service.GetSetmealsByID(ctx, req.(*paramModel.GetSetmealByIDRequest))
		},
		c.ShouldBindUri)
}
