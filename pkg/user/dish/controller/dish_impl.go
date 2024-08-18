package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/internal/utils/convert"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/dish/service"
)

type DishControllerImpl struct {
	service service.DishServiceInterface
}

func NewDishController(service service.DishServiceInterface) *DishControllerImpl {
	return &DishControllerImpl{service: service}
}

// GetDishByID 根据ID获取菜品
func (d *DishControllerImpl) GetDishByID(c *gin.Context) {
	categoryIdStr := c.Query("id")
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	categoryID, err := convert.StringToUint(categoryIdStr)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}

	dishes, apiError := d.service.GetDishByID(ctx, categoryID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, dishes)
}
