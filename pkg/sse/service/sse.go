package service

import (
	sseModel "sky-take-out-gin/pkg/sse/DTO"
)

// ClientChan 定义了客户端通道
type ClientChan chan sseModel.Message

// Client 定义了客户端结构
type Client struct {
	// participant 客户端参与者
	participant sseModel.Participant

	// Chan 客户端通道
	ch ClientChan
}

// SSEvent 定义了SSE事件结构
type SSEvent struct {
	// NewClient 客户端连接事件
	NewClient chan Client

	// CloseClient 客户端关闭事件
	CloseClient chan Client

	// Message 消息事件
	Message chan sseModel.Message

	// Clients 客户端列表
	Clients map[sseModel.Participant]ClientChan

	// Messages 消息列表
	Messages map[sseModel.Participant][]sseModel.Message
}

func NewSSEvent() *SSEvent {
	sse := &SSEvent{
		NewClient:   make(chan Client),
		CloseClient: make(chan Client),
		Message:     make(chan sseModel.Message),
		Clients:     make(map[sseModel.Participant]ClientChan),
		Messages:    make(map[sseModel.Participant][]sseModel.Message),
	}
	go sse.listen()
	return sse
}

// listen 监听事件
func (sse *SSEvent) listen() {
	for {
		select {
		case client := <-sse.NewClient: // 将客户端加入到客户端列表
			sse.Clients[client.participant] = client.ch
		case client := <-sse.CloseClient: // 将客户端从客户端列表中删除
			delete(sse.Clients, client.participant)
		case message := <-sse.Message:
			if client, ok := sse.Clients[message.To]; ok { // 如果消息的目标客户端存在，则直接发送消息
				client <- message
			}
			// 如果消息的目标客户端不存在，则将消息存储到消息列表中
			sse.Messages[message.To] = append(sse.Messages[message.To], message)
		}
	}
}

// AddClient 添加新的客户端
func (sse *SSEvent) AddClient(participant sseModel.Participant) ClientChan {
	clientChan := make(ClientChan)
	client := Client{
		participant: participant,
		ch:          clientChan,
	}
	sse.NewClient <- client
	return clientChan
}

// RemoveClient 移除客户端
func (sse *SSEvent) RemoveClient(participant sseModel.Participant) {
	client := Client{
		participant: participant,
	}
	sse.CloseClient <- client
}

// SendMessage 发送消息
func (sse *SSEvent) SendMessage(message sseModel.Message) {
	sse.Message <- message
}
