package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/code"
	"sky-take-out-gin/model"
	"time"
)

// BindRequest 绑定请求参数
// @Param c *gin.Context gin上下文
// @Param req interface{} 请求参数
// @Return *model.ApiError 错误信息
func BindRequest(c *gin.Context, req interface{}) *model.ApiError {
	if err := c.ShouldBind(req); err != nil {
		return &model.ApiError{
			Code: code.EmployeeBindParamError,
			Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
				code.EmployeeBindParamError,
				code.EmployeeBindParamError.Message(),
				err.Error()),
		}
	}
	return nil
}

// HandleRequest 处理请求的通用方法
// @Param c *gin.Context gin上下文
// @Param req interface{} 请求参数
// @Param serviceFunc func(ctx context.Context, req interface{}) (successResponse interface{}, err error) 需要调用的service层的方法
// @Param successRespType interface{} 成功响应类型
// @Return
func HandleRequest(c *gin.Context, req interface{}, serviceFunc func(ctx context.Context, req interface{}) (successResponse interface{}, err *model.ApiError)) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	resultChannel := make(chan interface{})
	go func() {
		if err := BindRequest(c, req); err != nil {
			resultChannel <- err
			return
		}

		resp, apiError := serviceFunc(ctx, req)
		if apiError != nil {
			resultChannel <- apiError
			return
		}

		resultChannel <- resp
		return
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			ResponseErrorWithCode(c, http.StatusRequestTimeout, code.ServerError)
			return
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
			return
		}
	case result := <-resultChannel:
		switch result.(type) {
		case *model.ApiError:
			ResponseErrorWithApiError(c, http.StatusBadRequest, result.(*model.ApiError))
			return
		default:
			ResponseSuccess(c, result)
			return
		}
	default:
		ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}
}
