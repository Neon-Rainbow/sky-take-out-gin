package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/set_meal/service"
	"strconv"
)

type SetMealControllerImpl struct {
	service service.SetMealServiceInterface
}

func NewSetMealController(service service.SetMealServiceInterface) SetMealControllerImpl {
	return SetMealControllerImpl{service: service}
}

func (controller SetMealControllerImpl) GetSetMealList(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	categoryIDStr := c.Query("category_id")
	if categoryIDStr == "" {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}

	setMealList, apiError := controller.service.GetSetMealList(ctx, categoryID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, apiError)
	}
	response.ResponseSuccess(c, setMealList)
}

func (controller SetMealControllerImpl) GetSetMealDetail(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	setMealIDStr := c.Param("id")
	if setMealIDStr == "" {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}

	setMealID, err := strconv.Atoi(setMealIDStr)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}

	setMealDetail, apiError := controller.service.GetSetMealDetail(ctx, setMealID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, apiError)
	}
	response.ResponseSuccess(c, setMealDetail)
}
