package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/shopping_cart/DTO"
	"sky-take-out-gin/pkg/user/shopping_cart/service"
)

type UserShoppingCartControllerImpl struct {
	service.UserShoppingCartServiceInterface
}

func NewUserShoppingCartControllerImpl(service service.UserShoppingCartServiceInterface) *UserShoppingCartControllerImpl {
	return &UserShoppingCartControllerImpl{service}
}

// GetShoppingCartList 获取购物车列表
func (controller *UserShoppingCartControllerImpl) GetShoppingCartList(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}

	userID := ctx.Value("userID").(uint)

	cartList, apiError := controller.UserShoppingCartServiceInterface.GetCartList(ctx, userID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, apiError)
		return
	}
	response.ResponseSuccess(c, cartList)
	return
}

// AddToShoppingCart 添加商品到购物车
func (controller *UserShoppingCartControllerImpl) AddToShoppingCart(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}

	userID := ctx.Value("userID").(uint)

	var req DTO.ShoppingCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}

	apiError := controller.UserShoppingCartServiceInterface.AddToCart(ctx, userID, &req)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

func (controller *UserShoppingCartControllerImpl) RemoveFromShoppingCart(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}

	userID := ctx.Value("userID").(uint)

	apiError := controller.UserShoppingCartServiceInterface.DeleteCartByUserID(ctx, userID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

func (controller *UserShoppingCartControllerImpl) SubOneCommodity(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}

	userID := ctx.Value("userID").(uint)

	var req DTO.ShoppingCartRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}

	apiError := controller.UserShoppingCartServiceInterface.DeleteCart(ctx, userID, &req)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}
