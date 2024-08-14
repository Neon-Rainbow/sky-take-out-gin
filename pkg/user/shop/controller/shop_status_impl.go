package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/shop/service"
)

type ShopStatusControllerImpl struct {
	service service.ShopStatusServiceInterface
}

func (controller *ShopStatusControllerImpl) GetShopStatus(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}
	status, apiError := controller.service.GetShopStatus(ctx)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, status)
}

// NewShopStatusControllerImpl 创建 ShopStatusControllerImpl 实例
func NewShopStatusControllerImpl(service service.ShopStatusServiceInterface) *ShopStatusControllerImpl {
	return &ShopStatusControllerImpl{service: service}
}
