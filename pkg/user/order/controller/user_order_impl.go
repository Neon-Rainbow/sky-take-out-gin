package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/internal/utils/convert"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/order/DTO"
	"sky-take-out-gin/pkg/user/order/service"
	"strconv"
)

type UserOrderControllerImpl struct {
	service service.UserOrderServiceInterface
}

// ReminderOrder 催单
// @Summary 催单
// @Tags 用户订单
// @Accept json
// @Produce json
// @Param orderID path int true "订单ID"
// @Success 200 {object} Response
// @Router /user/order/reminder/{orderID} [get]
func (controller *UserOrderControllerImpl) ReminderOrder(c *gin.Context) {
	orderID, err := convert.StringToUint(c.Param("orderID"))
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}
	userID := ctx.Value("userID").(uint)
	apiError := controller.service.ReminderOrder(ctx, userID, orderID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *UserOrderControllerImpl) RepetitionOrder(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserOrderControllerImpl) GetHistoryOrder(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}
	userID := ctx.Value("userID").(uint)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}
	orders, total, apiError := controller.service.GetHistoryOrder(ctx, userID, page, pageSize)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"total":   total,
		"records": orders,
	})
}

func (controller *UserOrderControllerImpl) CancelOrder(c *gin.Context) {
	orderID, err := convert.StringToUint(c.Param("orderID"))
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}

	var cancelReasonDTO DTO.CancelOrderRequestDTO
	if err := c.ShouldBindJSON(&cancelReasonDTO); err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}

	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}
	userID := ctx.Value("userID").(uint)
	apiError := controller.service.CancelOrder(ctx, userID, orderID, cancelReasonDTO.CancelReason)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *UserOrderControllerImpl) GetOrderDetail(c *gin.Context) {
	orderID, err := convert.StringToUint(c.Param("orderID"))
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}
	userID := ctx.Value("userID").(uint)
	order, apiError := controller.service.GetOrderDetail(ctx, userID, orderID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, order)
}

func (controller *UserOrderControllerImpl) SubmitOrder(c *gin.Context) {
	var submitOrderRequestDTO DTO.SubmitOrderRequestDTO
	if err := c.ShouldBindJSON(&submitOrderRequestDTO); err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}
	userID := ctx.Value("userID").(uint)
	apiError := controller.service.SubmitOrder(ctx, userID, &submitOrderRequestDTO)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *UserOrderControllerImpl) PayOrder(c *gin.Context) {
	var payOrderRequestDTO DTO.PayOrderRequestDTO
	if err := c.ShouldBindJSON(&payOrderRequestDTO); err != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, api_error.NewApiError(code.ParamError, err))
		return
	}
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}

	userID := ctx.Value("userID").(uint)
	apiError := controller.service.PayOrder(ctx, userID, payOrderRequestDTO.OrderID, payOrderRequestDTO.PayMethod)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func NewUserOrderController(service service.UserOrderServiceInterface) *UserOrderControllerImpl {
	return &UserOrderControllerImpl{service: service}
}
