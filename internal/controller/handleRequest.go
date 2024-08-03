package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/code"
	"sky-take-out-gin/model"
)

// HandleRequest 处理请求的通用方法
// @Param c *gin.Context gin上下文
// @Param req interface{} 请求参数
// @Param serviceFunc func(ctx context.Context, req interface{}) (successResponse interface{}, err error) 需要调用的service层的方法
// @Param bindFunc ...func(interface{}) error 绑定请求参数的方法,默认使用ShouldBind
// @Return
func HandleRequest(c *gin.Context,
	req interface{},
	serviceFunc func(ctx context.Context, req interface{}) (successResponse interface{}, err *model.ApiError),
	bindFunc ...func(interface{}) error) {
	//ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	//defer cancel()

	ctx := c.Request.Context()

	userID, exist := c.Get("user_id")
	if !exist {
		userID = 0
	}

	ctx = context.WithValue(ctx, "user_id", userID)

	username, exist := c.Get("username")
	if !exist {
		username = "guest"
	}

	username = username.(string)

	ctx = context.WithValue(ctx, "username", username)

	resultChannel := make(chan interface{})
	go func() {
		if len(bindFunc) == 0 {
			err := c.ShouldBind(req)
			if err != nil {
				resultChannel <- &model.ApiError{
					Code: code.ParamError,
					Msg:  err.Error(),
				}
				return
			}
		} else {
			for _, bindFunc := range bindFunc {
				if err := bindFunc(req); err != nil {
					resultChannel <- &model.ApiError{
						Code: code.ParamError,
						Msg:  err.Error(),
					}
					return
				}
			}
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
	}
	ResponseErrorWithMsg(c, http.StatusInternalServerError, code.ServerError, "未知错误")
}
