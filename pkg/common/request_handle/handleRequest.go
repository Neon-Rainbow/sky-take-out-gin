package request_handle

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/common/response"
)

// HandleRequest 处理请求的通用方法
// @Param c *gin.Context gin上下文
// @Param req interface{} 请求参数
// @Param serviceFunc func(ctx context.Context, req interface{}) (successResponse interface{}, err error) 需要调用的service层的方法
// @Param bindFunc ...func(interface{}) error 绑定请求参数的方法,默认使用ShouldBind
// @Return
func HandleRequest(c *gin.Context,
	req interface{},
	serviceFunc func(ctx context.Context, req interface{}) (successResponse interface{}, err *error2.ApiError),
	bindFunc ...func(interface{}) error) {
	//ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	//defer cancel()

	ctx, err := SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	resultChannel := make(chan interface{})
	go func() {
		if len(bindFunc) == 0 {
			err := c.ShouldBind(req)
			if err != nil {
				resultChannel <- &error2.ApiError{
					Code: code.ParamError,
					Msg:  err.Error(),
				}
				return
			}
		} else {
			for _, bindFunc := range bindFunc {
				if err := bindFunc(req); err != nil {
					resultChannel <- &error2.ApiError{
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
			response.ResponseErrorWithCode(c, http.StatusRequestTimeout, code.ServerError)
			zap.L().Error("请求超时",
				zap.String("path", c.Request.URL.Path),
			)
			return
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
			zap.L().Error("请求被取消",
				zap.String("path", c.Request.URL.Path),
			)
			return
		}
	case result := <-resultChannel:
		switch result.(type) {
		case *error2.ApiError:
			response.ResponseErrorWithApiError(c, http.StatusBadRequest, result.(*error2.ApiError))
			zap.L().Error("请求失败",
				zap.String("path", c.Request.URL.Path),
				zap.Any("error", result),
			)
			return
		default:
			response.ResponseSuccess(c, result)
			return
		}
	}
	response.ResponseErrorWithMsg(c, http.StatusInternalServerError, code.ServerError, "未知错误")
	zap.L().Error("未知错误",
		zap.String("path", c.Request.URL.Path),
	)
	return
}

// SetUserIDAndUsernameToContext 设置用户ID和用户名到上下文
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
