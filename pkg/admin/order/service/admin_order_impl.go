package service

import (
	"context"
	"fmt"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/admin/order/DTO"
	"sky-take-out-gin/pkg/admin/order/dao"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	sseModel "sky-take-out-gin/pkg/sse/DTO"
	sseRoute "sky-take-out-gin/pkg/sse/controller"
	"time"
)

type AdminOrderServiceImpl struct {
	dao dao.AdminOrderDaoInterface
}

func (service *AdminOrderServiceImpl) CancelOrder(ctx context.Context, orderID uint, cancelReason string) *api_error.ApiError {
	err := service.dao.CancelOrder(ctx, orderID, cancelReason)
	if err != nil {
		return api_error.NewApiError(code.CancelOrderError, err)
	}

	userID, err := service.dao.GetUserIDByOrderID(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.CancelOrderError, err)
	}

	sseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		To: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		Content: sseModel.Content{
			Type:      sseModel.CancelOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      fmt.Sprintf("商家取消了订单,原因：%v", cancelReason),
		},
	}
	sseRoute.GetSseEvent().SendMessage(*sseMessage)
	return nil
}

func (service *AdminOrderServiceImpl) GetOrderStatistics(ctx context.Context) (interface{}, *api_error.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service *AdminOrderServiceImpl) FinishOrder(ctx context.Context, orderID uint) *api_error.ApiError {
	err := service.dao.FinishOrder(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.FinishOrderError, err)
	}

	userID, err := service.dao.GetUserIDByOrderID(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.FinishOrderError, err)
	}

	sseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		To: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		Content: sseModel.Content{
			Type:      sseModel.FinishOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      "商家已完成订单",
		},
	}
	sseRoute.GetSseEvent().SendMessage(*sseMessage)
	return nil
}

func (service *AdminOrderServiceImpl) RejectOrder(ctx context.Context, orderID uint, rejectReason string) *api_error.ApiError {
	err := service.dao.RejectOrder(ctx, orderID, rejectReason)
	if err != nil {
		return api_error.NewApiError(code.RejectOrderError, err)
	}

	userID, err := service.dao.GetUserIDByOrderID(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.RejectOrderError, err)
	}

	sseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		To: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		Content: sseModel.Content{
			Type:      sseModel.RejectOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      fmt.Sprintf("商家拒绝了订单,原因：%v", rejectReason),
		},
	}
	sseRoute.GetSseEvent().SendMessage(*sseMessage)
	return nil
}

func (service *AdminOrderServiceImpl) ConfirmOrder(ctx context.Context, orderID uint) *api_error.ApiError {
	err := service.dao.ConfirmOrder(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.ConfirmOrderError, err)
	}

	userID, err := service.dao.GetUserIDByOrderID(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.ConfirmOrderError, err)
	}

	sseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		To: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		Content: sseModel.Content{
			Type:      sseModel.AcceptOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      "商家已确认订单",
		},
	}
	sseRoute.GetSseEvent().SendMessage(*sseMessage)
	return nil

}

func (service *AdminOrderServiceImpl) GetOrderByID(ctx context.Context, orderID uint) (*model.Order, *api_error.ApiError) {
	order, err := service.dao.GetOrderDetail(ctx, orderID)
	if err != nil {
		return nil, api_error.NewApiError(code.GetOrderDetailError, err)
	}
	return order, nil
}

func (service *AdminOrderServiceImpl) DeliveryOrder(ctx context.Context, orderID uint) *api_error.ApiError {
	err := service.dao.DeliveryOrder(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.DeliveryOrderError, err)
	}

	userID, err := service.dao.GetUserIDByOrderID(ctx, orderID)
	if err != nil {
		return api_error.NewApiError(code.DeliveryOrderError, err)
	}
	sseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		To: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		Content: sseModel.Content{
			Type:      sseModel.AcceptOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      "商家已发货",
		},
	}
	sseRoute.GetSseEvent().SendMessage(*sseMessage)
	return nil
}

func (service *AdminOrderServiceImpl) ConditionSearchOrder(ctx context.Context, params *DTO.QueryParams) ([]model.Order, *api_error.ApiError) {
	var orders []model.Order
	var err error
	orders, err = service.dao.ConditionSearchOrder(ctx, params)
	if err != nil {
		return nil, api_error.NewApiError(code.ConditionSearchOrderError, err)
	}
	return orders, nil
}

func NewAdminOrderService(dao dao.AdminOrderDaoInterface) *AdminOrderServiceImpl {
	return &AdminOrderServiceImpl{dao}
}
