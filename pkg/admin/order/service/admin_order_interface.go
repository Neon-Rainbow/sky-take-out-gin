package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/admin/order/DTO"
	"sky-take-out-gin/pkg/common/api_error"
)

type AdminOrderServiceInterface interface {
	CancelOrder(ctx context.Context, orderID uint, cancelReason string) *api_error.ApiError
	GetOrderStatistics(ctx context.Context) (interface{}, *api_error.ApiError)
	FinishOrder(ctx context.Context, orderID uint) *api_error.ApiError
	RejectOrder(ctx context.Context, orderID uint, rejectReason string) *api_error.ApiError
	ConfirmOrder(ctx context.Context, orderID uint) *api_error.ApiError
	GetOrderByID(ctx context.Context, orderID uint) (*model.Order, *api_error.ApiError)
	DeliveryOrder(ctx context.Context, orderID uint) *api_error.ApiError
	ConditionSearchOrder(ctx context.Context, params *DTO.QueryParams) ([]model.Order, *api_error.ApiError)
}
