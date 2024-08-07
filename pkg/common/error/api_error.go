package error

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
