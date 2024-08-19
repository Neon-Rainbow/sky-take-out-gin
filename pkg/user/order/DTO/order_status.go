package DTO

// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款
const (
	OrderStatusWaitPay     = 1
	OrderStatusWaitReceive = 2
	OrderStatusReceived    = 3
	OrderStatusDelivering  = 4
	OrderStatusCompleted   = 5
	OrderStatusCanceled    = 6
	OrderStatusRefunded    = 7
)

// 支付状态 0未支付 1已支付 2退款
const (
	OrderPayStatusUnpaid   = 0
	OrderPayStatusPaid     = 1
	OrderPayStatusRefunded = 2
)

const (
	WechatPay = 1
	AliPay    = 2
)
