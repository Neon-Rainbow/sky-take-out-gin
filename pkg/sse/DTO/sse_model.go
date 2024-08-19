package DTO

import "time"

// MessageType 定义了消息类型的枚举
type MessageType int

const (
	SubmitOrder MessageType = 1 // 提交订单
	AcceptOrder MessageType = 2 // 接受订单
	RejectOrder MessageType = 3 // 拒绝订单
	FinishOrder MessageType = 4 // 完成订单
	RemindOrder MessageType = 5 // 提醒订单
	CancelOrder MessageType = 6 // 取消订单
	PayOrder    MessageType = 7 // 支付订单
)

// ParticipantType 定义了参与者类型的枚举
type ParticipantType int

const (
	User     ParticipantType = 1
	Merchant ParticipantType = 2
)

// Message 定义了SSE传递的消息结构
type Message struct {
	// From 消息来源
	From Participant `json:"from"`

	// To 消息目标
	To Participant `json:"to"`

	// Content 消息内容
	Content Content `json:"content"`
}

// Participant 定义了一个加入SSE的对象
type Participant struct {
	// ID 对象ID
	ID uint `json:"id"`

	// Type 对象类型  1:用户 2:商家
	Type ParticipantType `json:"type"`
}

type Content struct {
	// Type 消息类型
	Type      MessageType `json:"type"`
	OrderID   uint        `json:"order_id"`
	TimeStamp int64       `json:"time_stamp"`
	Time      time.Time   `json:"time"`
	Text      interface{} `json:"text"`
}
