package request_handle

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/response"
)

type BindCommonFunc func(interface{}) error

// HandleRequest 处理请求的通用方法
// @Param c *gin.Context gin上下文
// @Param serviceFunc func(ctx context.Context, req *T) (*R, *api_error.ApiError) 调用的service层方法
// @Param bindFunc ...func(interface{}) error 绑定请求参数的方法,默认使用ShouldBind
// @Return
func HandleRequest[T any, R any](
	c *gin.Context,
	serviceFunc func(ctx context.Context, req *T) (*R, *api_error.ApiError),
	bindFunc ...BindCommonFunc) {

	ctx, err := SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithApiError(c, http.StatusInternalServerError, api_error.NewApiError(code.ServerError, err))
		return
	}

	var apiError *api_error.ApiError
	var req *T
	var resp *R

	if len(bindFunc) == 0 {
		err := c.ShouldBind(req)
		if err != nil {
			apiError = &api_error.ApiError{
				Code: code.ParamError,
				Msg:  err.Error(),
			}
			goto RESPONSE
		}
	} else {
		for _, bindFunc := range bindFunc {
			err := bindFunc(req)
			if err != nil {
				apiError = &api_error.ApiError{
					Code: code.ParamError,
					Msg:  err.Error(),
				}
				goto RESPONSE
			}
		}
	}

	resp, apiError = serviceFunc(ctx, req)
	if apiError != nil {
		goto RESPONSE
	}

RESPONSE:
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		zap.L().Error("请求失败",
			zap.String("path", c.Request.URL.Path),
			zap.Any("error", apiError),
		)
		return
	}
	response.ResponseSuccess(c, resp)

	return
}

// SetUserIDAndUsernameToContext 设置用户ID和用户名到上下文
// @Param c *gin.Context gin上下文
// @Return context.Context 上下文
// @Return error 错误
func SetUserIDAndUsernameToContext(c *gin.Context) (context.Context, error) {
	ctx := c.Request.Context()
	userID, exist := c.Get("userID")
	if !exist {
		userID = uint(0)
	} else {
		switch id := userID.(type) {
		case int:
			userID = uint(id)
		case int64:
			userID = uint(id)
		case uint:
			userID = id
		default:
			zap.L().Error("非法的用户ID类型", zap.Any("userID", userID))
			return nil, errors.New("非法的用户ID类型")
		}
	}
	ctx = context.WithValue(ctx, "userID", userID)

	username, exist := c.Get("username")
	if !exist {
		username = "guest"
	} else {
		if uname, ok := username.(string); ok {
			username = uname
		} else {
			zap.L().Error("非法的username类型", zap.Any("username", username))
			return nil, errors.New("非法的username类型")
		}
	}

	ctx = context.WithValue(ctx, "username", username)

	return ctx, nil
}
