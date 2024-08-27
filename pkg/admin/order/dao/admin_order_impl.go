package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/admin/order/DTO"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/database"
	"time"
)

type AdminOrderDaoImpl struct {
	db database.DatabaseInterface
}

func (dao *AdminOrderDaoImpl) CancelOrder(ctx context.Context, orderID uint, cancelReason string) error {
	var err error
	err = dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"status":        code.OrderStatusCanceled,
		"cancel_reason": cancelReason,
		"cancel_time":   time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *AdminOrderDaoImpl) GetOrderStatistics(ctx context.Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (dao *AdminOrderDaoImpl) FinishOrder(ctx context.Context, orderID uint) error {
	var err error
	err = dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("status", code.OrderStatusCompleted).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *AdminOrderDaoImpl) RejectOrder(ctx context.Context, orderID uint, rejectReason string) error {
	// 通过一次性更新多个字段，减少数据库交互次数
	err := dao.db.GetDB().WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"status":        code.OrderStatusCanceled,
			"reject_reason": rejectReason,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (dao *AdminOrderDaoImpl) ConfirmOrder(ctx context.Context, orderID uint) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("status", code.OrderStatusReceived).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *AdminOrderDaoImpl) GetOrderDetail(ctx context.Context, orderID uint) (*model.Order, error) {
	var order model.Order
	var err error
	err = dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Preload("OrderDetail").Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (dao *AdminOrderDaoImpl) DeliveryOrder(ctx context.Context, orderID uint) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("status", code.OrderStatusDelivering).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *AdminOrderDaoImpl) ConditionSearchOrder(ctx context.Context, queryParams *DTO.QueryParams) ([]model.Order, error) {
	var orders []model.Order
	var err error
	db := dao.db.GetDB().WithContext(ctx).Model(&model.Order{})

	if queryParams.OrderID != 0 {
		db = db.Where("id = ?", queryParams.OrderID)
	}

	if queryParams.Status != 0 {
		db = db.Where("status = ?", queryParams.Status)
	}

	if queryParams.BeginTime != "" && queryParams.EndTime != "" {
		db = db.Where("created_at BETWEEN ? AND ?", queryParams.BeginTime, queryParams.EndTime)
	}

	// 分页
	offset := (queryParams.Page - 1) * queryParams.PageSize
	db = db.Offset(offset).Limit(queryParams.PageSize)

	err = db.Preload("OrderDetail").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (dao *AdminOrderDaoImpl) GetUserIDByOrderID(ctx context.Context, orderID uint) (uint, error) {
	var order model.Order
	var err error
	err = dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return 0, err
	}
	return order.UserID, nil
}

func NewAdminOrderDaoImpl(db database.DatabaseInterface) *AdminOrderDaoImpl {
	return &AdminOrderDaoImpl{db: db}
}
