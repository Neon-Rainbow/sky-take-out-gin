package setmeal

import (
	"context"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/setmeal"
)

type SetmealServiceInterface interface {
	UpdateSetmeal(ctx context.Context, req *setmeal.UpdateSetmealRequest) (resp *setmeal.UpdateSetmealResponse, apiError *model.ApiError)
	GetSetmealPage(ctx context.Context, req *setmeal.GetSetmealsPageRequest) (resp *setmeal.GetSetmealsPageResponse, apiError *model.ApiError)
	ChangeSetmealStatus(ctx context.Context, req *setmeal.UpdateSetmealStatusRequest) (resp *setmeal.UpdateSetmealStatusResponse, apiError *model.ApiError)
	DeleteSetmeals(ctx context.Context, req *setmeal.DeleteSetmealsRequest) (resp *setmeal.DeleteSetmealsResponse, apiError *model.ApiError)
	CreateSetmeals(ctx context.Context, req *setmeal.AddSetmealRequest) (resp *setmeal.AddSetmealResponse, apiError *model.ApiError)
	GetSetmealsByID(ctx context.Context, req *setmeal.GetSetmealByIDRequest) (resp *setmeal.GetSetmealByIDResponse, apiError *model.ApiError)
}
