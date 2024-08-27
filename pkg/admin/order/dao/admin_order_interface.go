package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/admin/order/DTO"
)

type AdminOrderDaoInterface interface {
	CancelOrder(ctx context.Context, orderID uint, cancelReason string) error
	GetOrderStatistics(ctx context.Context) interface{}
	FinishOrder(ctx context.Context, orderID uint) error
	RejectOrder(ctx context.Context, orderID uint, rejectReason string) error
	ConfirmOrder(ctx context.Context, orderID uint) error
	GetOrderDetail(ctx context.Context, orderID uint) (*model.Order, error)
	DeliveryOrder(ctx context.Context, orderID uint) error
	ConditionSearchOrder(ctx context.Context, queryParams *DTO.QueryParams) ([]model.Order, error)
	GetUserIDByOrderID(ctx context.Context, orderID uint) (uint, error)
}
