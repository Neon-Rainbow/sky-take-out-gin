package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type UserOrderDaoImpl struct {
	db database.DatabaseInterface
}

func (dao *UserOrderDaoImpl) CreateOrder(ctx context.Context, order *model.Order) error {
	err := dao.db.GetDB().Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserOrderDaoImpl) UpdateOrder(ctx context.Context, order *model.Order) error {
	err := dao.db.GetDB().Save(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserOrderDaoImpl) GetOrderByID(ctx context.Context, orderID uint) (*model.Order, error) {
	order := &model.Order{}
	err := dao.db.GetDB().WithContext(ctx).Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (dao *UserOrderDaoImpl) GetOrderPage(ctx context.Context, userID uint, page int, size int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64
	var err error

	err = dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.db.GetDB().WithContext(ctx).Where("user_id = ?", userID).Offset((page - 1) * size).Limit(size).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (dao *UserOrderDaoImpl) UpdateOrderStatus(ctx context.Context, orderID uint, status int) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserOrderDaoImpl) UpdateOrderPayStatus(ctx context.Context, orderID uint, payStatus int, payMethod int) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("pay_status", payStatus).Error
	if err != nil {
		return err
	}
	err = dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("pay_method", payMethod).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserOrderDaoImpl) UpdateOrderDeliveryStatus(ctx context.Context, orderID uint, deliveryStatus int) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update("delivery_status", deliveryStatus).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateOrderColumn 更新订单的某个字段
func (dao *UserOrderDaoImpl) UpdateOrderColumn(ctx context.Context, orderID uint, column string, value interface{}) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Update(column, value).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateOrderColumns 更新订单的多个字段
func (dao *UserOrderDaoImpl) UpdateOrderColumns(ctx context.Context, orderID uint, columns map[string]interface{}) error {
	err := dao.db.GetDB().WithContext(ctx).Model(&model.Order{}).Where("id = ?", orderID).Updates(columns).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserOrderDaoImpl(db database.DatabaseInterface) *UserOrderDaoImpl {
	return &UserOrderDaoImpl{db: db}
}
