package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/admin/order/DTO"
	"sky-take-out-gin/pkg/admin/order/service"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/response"
)

type AdminOrderControllerImpl struct {
	service service.AdminOrderServiceInterface
}

func (controller *AdminOrderControllerImpl) CancelOrder(c *gin.Context) {
	var cancelOrderRequestDTO DTO.CancelOrderRequestDTO
	if err := c.ShouldBindBodyWithJSON(&cancelOrderRequestDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	apiError := controller.service.CancelOrder(c, cancelOrderRequestDTO.OrderID, cancelOrderRequestDTO.CancelReason)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *AdminOrderControllerImpl) GetOrderStatistics(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *AdminOrderControllerImpl) FinishOrder(c *gin.Context) {
	var finishOrderRequestDTO DTO.OrderIDDTO
	if err := c.ShouldBindUri(&finishOrderRequestDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	apiError := controller.service.FinishOrder(c, finishOrderRequestDTO.OrderID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *AdminOrderControllerImpl) RejectOrder(c *gin.Context) {
	var rejectOrderRequestDTO DTO.RejectOrderRequestDTO
	if err := c.ShouldBindBodyWithJSON(&rejectOrderRequestDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	apiError := controller.service.RejectOrder(c, rejectOrderRequestDTO.OrderID, rejectOrderRequestDTO.RejectReason)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *AdminOrderControllerImpl) ConfirmOrder(c *gin.Context) {
	var orderIDDTO DTO.OrderIDDTO
	if err := c.ShouldBindBodyWithJSON(&orderIDDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	apiError := controller.service.ConfirmOrder(c, orderIDDTO.OrderID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *AdminOrderControllerImpl) GetOrderByID(c *gin.Context) {
	var orderIDDTO DTO.OrderIDDTO
	if err := c.ShouldBindUri(&orderIDDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	order, apiError := controller.service.GetOrderByID(c, orderIDDTO.OrderID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, order)
}

func (controller *AdminOrderControllerImpl) DeliveryOrder(c *gin.Context) {
	var orderIDDTO DTO.OrderIDDTO
	if err := c.ShouldBindUri(&orderIDDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	apiError := controller.service.DeliveryOrder(c, orderIDDTO.OrderID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *AdminOrderControllerImpl) ConditionSearchOrder(c *gin.Context) {
	var queryParams DTO.QueryParams
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}
	orders, apiError := controller.service.ConditionSearchOrder(c, &queryParams)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, orders)
}

func NewAdminOrderController(service service.AdminOrderServiceInterface) *AdminOrderControllerImpl {
	return &AdminOrderControllerImpl{service: service}
}
