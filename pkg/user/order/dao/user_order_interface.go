package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type UserOrderDaoInterface interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
	GetOrderByID(ctx context.Context, orderID uint) (*model.Order, error)
	GetOrderPage(ctx context.Context, userID uint, page int, size int) (orders []model.Order, total int64, err error)
	UpdateOrderStatus(ctx context.Context, orderID uint, status int) error
	UpdateOrderPayStatus(ctx context.Context, orderID uint, payStatus int, payMethod int) error
	UpdateOrderDeliveryStatus(ctx context.Context, orderID uint, deliveryStatus int) error
	UpdateOrderColumn(ctx context.Context, orderID uint, column string, value interface{}) error
	UpdateOrderColumns(ctx context.Context, orderID uint, columns map[string]interface{}) error
}
