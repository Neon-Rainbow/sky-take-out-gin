package DTO

// MessageType 定义了消息类型的枚举
type MessageType int

const (
	UserToMerchant MessageType = 1
	MerchantToUser MessageType = 2
)

// ParticipantType 定义了参与者类型的枚举
type ParticipantType int

const (
	User     ParticipantType = 1
	Merchant ParticipantType = 2
)

// Message 定义了SSE传递的消息结构
type Message struct {
	// Type 消息类型  1:用户->商家,2:商家->用户
	Type MessageType `json:"type"`

	// From 消息来源
	From Participant `json:"from"`

	// To 消息目标
	To Participant `json:"to"`

	// Content 消息内容
	Content interface{} `json:"content"`
}

// Participant 定义了一个加入SSE的对象
type Participant struct {
	// ID 对象ID
	ID int64 `json:"id"`

	// Type 对象类型  1:用户 2:商家
	Type ParticipantType `json:"type"`
}
