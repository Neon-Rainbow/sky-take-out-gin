package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	apiErrorModel "sky-take-out-gin/pkg/common/error"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/user/category/DTO"
	"sky-take-out-gin/pkg/user/category/service"
)

type CategoryControllerImpl struct {
	service.CategoryServiceInterface
}

func (controller CategoryControllerImpl) GetCategoryList(c *gin.Context) {
	req := DTO.CategoryRequestDTO{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (successResponse interface{}, apiError *apiErrorModel.ApiError) {
			return controller.CategoryServiceInterface.GetCategoryList(ctx, req.(*DTO.CategoryRequestDTO))
		},
		c.ShouldBindQuery,
	)
}

func NewCategoryController(service service.CategoryServiceInterface) CategoryControllerImpl {
	return CategoryControllerImpl{service}
}
