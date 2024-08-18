package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/category/DTO"
	"sky-take-out-gin/pkg/user/category/service"
)

type CategoryControllerImpl struct {
	service.CategoryServiceInterface
}

func (controller CategoryControllerImpl) GetCategoryList(c *gin.Context) {
	req := DTO.CategoryRequestDTO{}

	if err := c.ShouldBindQuery(&req); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	resp, apiError := controller.CategoryServiceInterface.GetCategoryList(c.Request.Context(), &req)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, resp.Categories)
}

func NewCategoryController(service service.CategoryServiceInterface) CategoryControllerImpl {
	return CategoryControllerImpl{service}
}
