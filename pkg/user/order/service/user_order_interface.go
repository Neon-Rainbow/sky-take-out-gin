package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/user/order/DTO"
)

type UserOrderServiceInterface interface {
	ReminderOrder(ctx context.Context, userID uint, orderID uint) *api_error.ApiError
	RepetitionOrder(ctx context.Context, userID uint, orderID uint) *api_error.ApiError
	GetHistoryOrder(ctx context.Context, userID uint, page int, size int) ([]model.Order, int64, *api_error.ApiError)
	CancelOrder(ctx context.Context, userID uint, orderID uint, cancelReason string) *api_error.ApiError
	GetOrderDetail(ctx context.Context, userID uint, orderID uint) (*model.Order, *api_error.ApiError)
	SubmitOrder(ctx context.Context, userID uint, orderRequestDTO *DTO.SubmitOrderRequestDTO) *api_error.ApiError
	PayOrder(ctx context.Context, userID uint, orderID uint, payMethod int) *api_error.ApiError
}
