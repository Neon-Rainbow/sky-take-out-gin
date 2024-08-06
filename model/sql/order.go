package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Number                string     `gorm:"type:varchar(50);comment:'订单号'"`
	Status                int        `gorm:"default:1;comment:'订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款'"`
	UserID                int64      `gorm:"not null;comment:'下单用户'"`
	AddressBookID         int64      `gorm:"not null;comment:'地址id'"`
	OrderTime             time.Time  `gorm:"not null;comment:'下单时间'"`
	CheckoutTime          *time.Time `gorm:"comment:'结账时间'"`
	PayMethod             int        `gorm:"default:1;comment:'支付方式 1微信,2支付宝'"`
	PayStatus             int        `gorm:"default:0;comment:'支付状态 0未支付 1已支付 2退款'"`
	Amount                float64    `gorm:"type:decimal(10,2);not null;comment:'实收金额'"`
	Remark                string     `gorm:"type:varchar(100);comment:'备注'"`
	Phone                 string     `gorm:"type:varchar(11);comment:'手机号'"`
	Address               string     `gorm:"type:varchar(255);comment:'地址'"`
	UserName              string     `gorm:"type:varchar(32);comment:'用户名称'"`
	Consignee             string     `gorm:"type:varchar(32);comment:'收货人'"`
	CancelReason          string     `gorm:"type:varchar(255);comment:'订单取消原因'"`
	RejectionReason       string     `gorm:"type:varchar(255);comment:'订单拒绝原因'"`
	CancelTime            *time.Time `gorm:"comment:'订单取消时间'"`
	EstimatedDeliveryTime *time.Time `gorm:"comment:'预计送达时间'"`
	DeliveryStatus        int        `gorm:"default:1;comment:'配送状态  1立即送出  0选择具体时间'"`
	DeliveryTime          *time.Time `gorm:"comment:'送达时间'"`
	PackAmount            int        `gorm:"comment:'打包费'"`
	TablewareNumber       int        `gorm:"comment:'餐具数量'"`
	TablewareStatus       int        `gorm:"default:1;comment:'餐具数量状态  1按餐量提供  0选择具体数量'"`
}
