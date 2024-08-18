package service

import (
	"context"
	apiErrorModel "sky-take-out-gin/pkg/common/api_error"
)

type UserServiceInterface interface {
	Login(ctx context.Context, username string, password string) (accessToken string, refreshToken string, userID uint, apiError *apiErrorModel.ApiError)
	Logout(ctx context.Context) *apiErrorModel.ApiError
	RefreshToken(ctx context.Context, refreshToken string) (accessToken string, newRefreshToken string, apiError *apiErrorModel.ApiError)
	Register(ctx context.Context, username string, password string) (userID uint, apiError *apiErrorModel.ApiError)
}
