package sseRoute

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	sseModel "sky-take-out-gin/pkg/sse/DTO"
	"sky-take-out-gin/pkg/sse/service"
	"strconv"
	"time"
)

var sseEvent = service.NewSSEvent()

func GetSseEvent() *service.SSEvent {
	return sseEvent
}

// SSEHandler 定义了SSE事件处理函数
// @Summary SSE事件处理函数
// @Description 通过SSE协议向客户端推送消息,该函数与客户端建立长连接，当有消息到达时，向客户端推送消息
func SSEHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	idTypeStr := c.Query("type")
	idType, err := strconv.ParseInt(idTypeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "非法type",
		})
	}
	var t sseModel.ParticipantType
	if idType == 1 {
		t = sseModel.User
	} else {
		t = sseModel.Merchant
	}

	participant := sseModel.Participant{
		ID:   uint(id),
		Type: t,
	}

	clientChan := sseEvent.AddClient(participant)

	defer sseEvent.RemoveClient(participant)

	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-clientChan:
			if ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		case <-time.After(50 * time.Second):
			// 如果长时间没有消息，则发送一个心跳消息，保持连接
			c.SSEvent("heartbeat", "keepalive")
			return true
		}
	})
}

// SendMessageHandler 定义了发送消息处理函数
func SendMessageHandler(c *gin.Context) {
	var msg sseModel.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sseEvent.SendMessage(msg)
	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
}
