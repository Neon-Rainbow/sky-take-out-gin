package DTO

import "time"

// SubmitOrderRequestDTO 提交订单请求DTO
type SubmitOrderRequestDTO struct {
	AddressBookID         uint      `json:"address_book_ID" binding:"required"`
	Amount                int       `json:"amount" binding:"required"`
	DeliveryStatus        int       `json:"delivery_status" binding:"required"`
	EstimatedDeliveryTime time.Time `json:"estimated_delivery_time" binding:"required"`
	PackAmount            int       `json:"pack_amount" binding:"required"`
	PayMethod             int       `json:"pay_method" binding:"required"`
	Remark                string    `json:"remark"`
	TablewareAmount       int       `json:"tableware_amount" binding:"required"`
	TablewareStatus       int       `json:"tableware_status" binding:"required"`
}

type PayOrderRequestDTO struct {
	OrderID   uint `json:"order_id" binding:"required"`
	PayMethod int  `json:"pay_method" binding:"required"`
}

type CancelOrderRequestDTO struct {
	CancelReason string `json:"cancel_reason" binding:"required"`
}
