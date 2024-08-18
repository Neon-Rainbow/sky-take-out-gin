package api_error

import (
	"fmt"
	"sky-take-out-gin/pkg/common/code"
)

type ApiError struct {
	Code code.ResponseCode `json:"code"`
	Msg  string            `json:"msg"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("ApiError: code=%v, msg=%v", e.Code, e.Msg)
}

// NewApiError 创建一个新的ApiError
// @Param code code.ResponseCode 错误码
// @Param args ...interface{} 错误信息，可以是string或error类型
// @Return *ApiError
func NewApiError(code code.ResponseCode, args ...interface{}) *ApiError {
	var msg string

	if len(args) == 0 {
		msg = code.Message()
	} else {
		switch v := args[0].(type) {
		case string:
			msg = v
		case error:
			msg = v.Error()
		default:
			msg = code.Message() // 如果传入的类型不是string或error，则使用默认消息
		}
	}

	return &ApiError{
		Code: code,
		Msg:  msg,
	}
}
