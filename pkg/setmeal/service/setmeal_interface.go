package service

import (
	"context"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/setmeal/DTO"
)

type SetmealServiceInterface interface {
	UpdateSetmeal(ctx context.Context, req *DTO.UpdateSetmealRequest) (resp *DTO.UpdateSetmealResponse, apiError *error2.ApiError)
	GetSetmealPage(ctx context.Context, req *DTO.GetSetmealsPageRequest) (resp *DTO.GetSetmealsPageResponse, apiError *error2.ApiError)
	ChangeSetmealStatus(ctx context.Context, req *DTO.UpdateSetmealStatusRequest) (resp *DTO.UpdateSetmealStatusResponse, apiError *error2.ApiError)
	DeleteSetmeals(ctx context.Context, req *DTO.DeleteSetmealsRequest) (resp *DTO.DeleteSetmealsResponse, apiError *error2.ApiError)
	CreateSetmeals(ctx context.Context, req *DTO.AddSetmealRequest) (resp *DTO.AddSetmealResponse, apiError *error2.ApiError)
	GetSetmealsByID(ctx context.Context, req *DTO.GetSetmealByIDRequest) (resp *DTO.GetSetmealByIDResponse, apiError *error2.ApiError)
}
