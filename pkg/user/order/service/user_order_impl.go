package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	sseModel "sky-take-out-gin/pkg/sse/DTO"
	sseRoute "sky-take-out-gin/pkg/sse/controller"
	dao2 "sky-take-out-gin/pkg/user/address_book/dao"
	"sky-take-out-gin/pkg/user/order/DTO"
	"sky-take-out-gin/pkg/user/order/dao"
	dao4 "sky-take-out-gin/pkg/user/shopping_cart/dao"
	dao3 "sky-take-out-gin/pkg/user/user/dao"
	"time"
)

type UserOrderServiceImpl struct {
	userOrderDao    dao.UserOrderDaoInterface
	userAddressDao  dao2.AddressBookDaoInterface
	userDao         dao3.UserDaoInterface
	shoppingCartDao dao4.UserShoppingCartDaoCacheInterface
}

func (service *UserOrderServiceImpl) ReminderOrder(ctx context.Context, userID uint, orderID uint) *api_error.ApiError {
	SseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		To: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		Content: sseModel.Content{
			Type:      sseModel.RemindOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      "用户催单",
		},
	}
	sseRoute.GetSseEvent().SendMessage(*SseMessage) // 发送消息
	return nil
}

func (service *UserOrderServiceImpl) RepetitionOrder(ctx context.Context, userID uint, orderID uint) *api_error.ApiError {
	//TODO implement me
	panic("implement me")
}

func (service *UserOrderServiceImpl) GetHistoryOrder(ctx context.Context, userID uint, page int, size int) ([]model.Order, int64, *api_error.ApiError) {
	orders, total, err := service.userOrderDao.GetOrderPage(ctx, userID, page, size)
	if err != nil {
		return nil, 0, api_error.NewApiError(code.GetHistoryOrderError, err)
	}
	return orders, total, nil
}

func (service *UserOrderServiceImpl) CancelOrder(ctx context.Context, userID uint, orderID uint, cancelReason string) *api_error.ApiError {

	var updateColumns = make(map[string]interface{})
	updateColumns["status"] = DTO.OrderStatusCanceled
	updateColumns["cancel_reason"] = cancelReason
	updateColumns["cancel_time"] = time.Now()

	err := service.userOrderDao.UpdateOrderColumns(ctx, orderID, updateColumns)
	if err != nil {
		return api_error.NewApiError(code.CancelOrderError, err)
	}

	SseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		To: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		Content: sseModel.Content{
			Type:      sseModel.CancelOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      fmt.Sprintf("用户取消了订单,原因：%v", cancelReason),
		},
	}
	sseRoute.GetSseEvent().SendMessage(*SseMessage) // 发送消息
	return nil
}

func (service *UserOrderServiceImpl) GetOrderDetail(ctx context.Context, userID uint, orderID uint) (*model.Order, *api_error.ApiError) {
	order, err := service.userOrderDao.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, api_error.NewApiError(code.GetOrderDetailError, err)
	}
	return order, nil
}

func (service *UserOrderServiceImpl) SubmitOrder(ctx context.Context, userID uint, orderRequestDTO *DTO.SubmitOrderRequestDTO) *api_error.ApiError {
	order := &model.Order{}
	address := &model.AddressBook{}
	user := &model.User{}

	var err error

	err = copier.CopyWithOption(order, orderRequestDTO, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return api_error.NewApiError(code.SubmitOrderError, err)
	}

	order.OrderTime = time.Now()
	order.UserID = userID

	user, err = service.userDao.GetUserByID(ctx, userID)
	if err != nil {
		return api_error.NewApiError(code.SubmitOrderError, err)
	}
	order.UserName = user.Username

	address, err = service.userAddressDao.GetAddressByID(ctx, order.AddressBookID)
	if err != nil {
		return api_error.NewApiError(code.SubmitOrderError, err)
	}

	order.Address = address.GetDetailAddress()
	order.Consignee = address.Consignee
	order.Phone = address.Phone

	lists, err := service.shoppingCartDao.GetCartList(ctx, userID)
	if err != nil {
		return api_error.NewApiError(code.SubmitOrderError, err)
	}
	for _, v := range lists {
		tempOrderDetail := model.OrderDetail{}
		err = copier.CopyWithOption(&tempOrderDetail, &v, copier.Option{IgnoreEmpty: true})
		if err != nil {
			return api_error.NewApiError(code.SubmitOrderError, err)
		}
		order.OrderDetail = append(order.OrderDetail, tempOrderDetail)
	}

	err = service.userOrderDao.CreateOrder(ctx, order)
	if err != nil {
		return api_error.NewApiError(code.SubmitOrderError, err)
	}

	SseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		To: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		Content: sseModel.Content{
			Type:      sseModel.SubmitOrder,
			OrderID:   order.ID,
			TimeStamp: time.Now().Unix(),
			Time:      time.Now(),
			Text:      "用户提交了订单",
		},
	}
	sseRoute.GetSseEvent().SendMessage(*SseMessage) // 发送消息
	return nil
}

func (service *UserOrderServiceImpl) PayOrder(ctx context.Context, userID uint, orderID uint, payMethod int) *api_error.ApiError {
	err := service.userOrderDao.UpdateOrderPayStatus(ctx, orderID, DTO.OrderPayStatusPaid, payMethod)
	if err != nil {
		return api_error.NewApiError(code.PayOrderError, err)
	}
	SseMessage := &sseModel.Message{
		From: sseModel.Participant{
			ID:   userID,
			Type: sseModel.User,
		},
		To: sseModel.Participant{
			ID:   1,
			Type: sseModel.Merchant,
		},
		Content: sseModel.Content{
			Type:      sseModel.PayOrder,
			OrderID:   orderID,
			TimeStamp: time.Now().Unix(),
			Text:      "用户支付了订单",
		},
	}
	sseRoute.GetSseEvent().SendMessage(*SseMessage) // 发送消息
	return nil
}

func NewUserOrderServiceImpl(
	userOrderDao dao.UserOrderDaoInterface,
	userAddressDao dao2.AddressBookDaoInterface,
	userInfoDao dao3.UserDaoInterface,
	shoppingCartDao dao4.UserShoppingCartDaoCacheInterface) *UserOrderServiceImpl {
	return &UserOrderServiceImpl{
		userOrderDao,
		userAddressDao,
		userInfoDao,
		shoppingCartDao}
}
