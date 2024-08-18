package service

import (
	"context"
	"sky-take-out-gin/internal/utils/encrypt"
	"sky-take-out-gin/pkg/common/JWT"
	error2 "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/user/user/dao"
	"time"
)

type UserServiceImpl struct {
	dao.UserDaoInterface
	dao.UserTokenDAOInterface
}

func (service *UserServiceImpl) Login(ctx context.Context, username string, password string) (accessToken string, refreshToken string, userID uint, apiError *error2.ApiError) {
	userID, hashedPassword, err := service.UserDaoInterface.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", 0, &error2.ApiError{
			Code: code.UserLoginError,
			Msg:  err.Error(),
		}
	}

	if encrypt.EncryptPassword(password) != hashedPassword {
		return "", "", 0, &error2.ApiError{
			Code: code.UserLoginError,
			Msg:  "password not match",
		}
	}
	accessToken, refreshToken, err = JWT.GenerateToken(username, userID, "user")
	if err != nil {
		return "", "", 0, &error2.ApiError{
			Code: code.UserLoginError,
			Msg:  err.Error(),
		}
	}
	err = service.UserTokenDAOInterface.SaveToken(ctx, userID, accessToken, refreshToken, time.Hour*24)
	if err != nil {
		return "", "", 0, &error2.ApiError{
			Code: code.UserLoginError,
			Msg:  err.Error(),
		}
	}
	return accessToken, refreshToken, userID, nil
}

func (service *UserServiceImpl) Logout(ctx context.Context) *error2.ApiError {
	err := service.UserTokenDAOInterface.DeleteTokens(ctx, ctx.Value("userID").(uint))
	if err != nil {
		return &error2.ApiError{
			Code: code.UserLogoutError,
			Msg:  err.Error(),
		}
	}
	return nil
}

func (service *UserServiceImpl) RefreshToken(ctx context.Context, refreshToken string) (accessToken string, newRefreshToken string, apiError *error2.ApiError) {
	myClaims, err := JWT.ParseToken(refreshToken)
	if err != nil {
		return "", "", &error2.ApiError{
			Code: code.UserRefreshTokenError,
			Msg:  err.Error(),
		}
	}
	userID := myClaims.UserID

	valid, err := service.UserTokenDAOInterface.ValidateRefreshToken(ctx, userID, refreshToken)
	if err != nil {
		return "", "", &error2.ApiError{
			Code: code.UserRefreshTokenError,
			Msg:  err.Error(),
		}
	}
	if !valid {
		return "", "", &error2.ApiError{
			Code: code.UserRefreshTokenError,
			Msg:  "非法的refresh token,可能是因为已经被使用过了,或者已经过期,请重新登录",
		}
	}
	accessToken, newRefreshToken, err = JWT.GenerateToken("", userID, "user")
	if err != nil {
		return "", "", &error2.ApiError{
			Code: code.UserRefreshTokenError,
			Msg:  err.Error(),
		}
	}
	err = service.UserTokenDAOInterface.SaveToken(ctx, userID, accessToken, newRefreshToken, time.Hour*24)
	if err != nil {
		return "", "", &error2.ApiError{
			Code: code.UserRefreshTokenError,
			Msg:  err.Error(),
		}
	}
	return accessToken, newRefreshToken, nil
}

// Register 注册
func (service *UserServiceImpl) Register(ctx context.Context, username string, password string) (userID uint, apiError *error2.ApiError) {
	// 检查用户名是否存在
	_, _, err := service.UserDaoInterface.GetUserByUsername(ctx, username)
	if err == nil {
		return 0, &error2.ApiError{
			Code: code.UserRegisterError,
			Msg:  "username already exists",
		}
	}

	hashedPassword := encrypt.EncryptPassword(password)
	userID, err = service.UserDaoInterface.CreateUser(ctx, username, hashedPassword)
	if err != nil {
		return 0, &error2.ApiError{
			Code: code.UserRegisterError,
			Msg:  err.Error(),
		}
	}
	return userID, nil
}

// NewUserServiceImpl 创建 UserServiceImpl 实例
func NewUserServiceImpl(daoImpl dao.UserDaoInterface, daoTokenImpl dao.UserTokenDAOInterface) *UserServiceImpl {
	return &UserServiceImpl{
		UserDaoInterface:      daoImpl,
		UserTokenDAOInterface: daoTokenImpl,
	}
}
