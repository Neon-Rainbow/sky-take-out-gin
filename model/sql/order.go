package model

import "time"

type Order struct {
	ID                    int64     `json:"id" gorm:"primary_key"`
	Number                string    `json:"number"`
	Status                int       `json:"status"`
	UserID                int64     `json:"user_id"`
	AddressBookID         int64     `json:"address_book_id"`
	OrderTime             time.Time `json:"order_time"`
	CheckoutTime          time.Time `json:"checkout_time"`
	PayMethod             int       `json:"pay_method"`
	PayStatus             int       `json:"pay_status"`
	Amount                float64   `json:"amount"`
	Remark                string    `json:"remark"`
	Phone                 string    `json:"phone"`
	Address               string    `json:"address"`
	UserName              string    `json:"user_name"`
	Consignee             string    `json:"consignee"`
	CancelReason          string    `json:"cancel_reason"`
	RejectionReason       string    `json:"rejection_reason"`
	CancelTime            time.Time `json:"cancel_time"`
	EstimatedDeliveryTime time.Time `json:"estimated_delivery_time"`
	DeliveryStatus        int       `json:"delivery_status"`
	DeliveryTime          time.Time `json:"delivery_time"`
	PackAmount            int       `json:"pack_amount"`
	TablewareNumber       int       `json:"tableware_number"`
	TablewareStatus       int       `json:"tableware_status"`
}

func (Order) TableName() string {
	return "order"
}
