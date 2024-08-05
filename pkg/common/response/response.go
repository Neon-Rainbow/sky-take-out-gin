package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
)

// Response 用于返回响应
type Response struct {
	Code code.ResponseCode `json:"code"`
	Msg  string            `json:"msg"`
	Data interface{}       `json:"data"`
}

// ResponseSuccess 用于返回成功的响应
// @param c gin.Context
// @param data interface{}
// @return
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 00000,
		Msg:  "success",
		Data: data,
	})
}

// ResponseErrorWithCode 用于返回错误的响应
// @param c gin.Context
// @param httpStatus int
// @param code code.ResponseCode
// @return
func ResponseErrorWithCode(c *gin.Context, httpStatus int, code code.ResponseCode) {
	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  code.Message(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 用于返回错误的响应
// @param c gin.Context
// @param httpStatus int
// @param code code.ResponseCode
// @param msg string
// @return
func ResponseErrorWithMsg(c *gin.Context, httpStatus int, code code.ResponseCode, msg string) {
	if msg == "" {
		msg = code.Message()
	}
	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseErrorWithApiError 用于返回错误的响应
// @param c gin.Context
// @param httpStatus int
// @param apiError *model.ApiError
// @return
func ResponseErrorWithApiError(c *gin.Context, httpStatus int, apiError *error2.ApiError) {
	var msg string
	if apiError.Msg == "" {
		msg = apiError.Code.Message()
	} else {
		msg = apiError.Msg
	}
	c.JSON(httpStatus, Response{
		Code: apiError.Code,
		Msg:  msg,
		Data: nil,
	})
}
