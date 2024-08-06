package sseRoute

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	sseModel "sky-take-out-gin/pkg/sse/DTO"
	"sky-take-out-gin/pkg/sse/service"
	"strconv"
)

var sseEvent = service.NewSSEvent()

// SSEHandler 定义了SSE事件处理函数
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
		ID:   id,
		Type: t,
	}

	clientChan := sseEvent.AddClient(participant)

	defer sseEvent.RemoveClient(participant)

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
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
