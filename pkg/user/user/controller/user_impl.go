package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/user/DTO"
	"sky-take-out-gin/pkg/user/user/service"
)

type UserLoginControllerImpl struct {
	service service.UserServiceInterface
}

func NewUserLoginController(service service.UserServiceInterface) UserLoginControllerInterface {
	return &UserLoginControllerImpl{service: service}
}

func (controller *UserLoginControllerImpl) Login(c *gin.Context) {
	var UserLoginRequestDTO DTO.UserLoginRequestDTO
	if err := c.ShouldBindJSON(&UserLoginRequestDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.UserLoginError)
		return
	}
	accessToken, refreshToken, userID, apiError := controller.service.Login(c.Request.Context(), UserLoginRequestDTO.Username, UserLoginRequestDTO.Password)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user_id":       userID,
	})
}

func (controller *UserLoginControllerImpl) Logout(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.UserLogoutError)
		return
	}
	apiError := controller.service.Logout(ctx)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller *UserLoginControllerImpl) Register(c *gin.Context) {
	var UserRegisterRequestDTO DTO.UserRegisterRequestDTO
	if err := c.ShouldBindJSON(&UserRegisterRequestDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.UserRegisterError)
		return
	}
	userID, apiError := controller.service.Register(c.Request.Context(), UserRegisterRequestDTO.Username, UserRegisterRequestDTO.Password)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"user_id": userID,
	})
}

func (controller *UserLoginControllerImpl) RefreshToken(c *gin.Context) {
	var UserRefreshTokenRequestDTO DTO.UserRefreshTokenRequestDTO
	if err := c.ShouldBindJSON(&UserRefreshTokenRequestDTO); err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.UserRefreshTokenError)
		return
	}
	accessToken, refreshToken, apiError := controller.service.RefreshToken(c.Request.Context(), UserRefreshTokenRequestDTO.RefreshToken)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
