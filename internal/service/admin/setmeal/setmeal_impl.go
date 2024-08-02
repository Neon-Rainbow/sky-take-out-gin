package setmeal

import (
	"context"
	setmealDAO "sky-take-out-gin/internal/dao/admin/setmeal"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/setmeal"
)

type SetmealServiceImpl struct {
	dao setmealDAO.SetmealDAOInterface
}

func (service SetmealServiceImpl) UpdateSetmeal(ctx context.Context, req *setmeal.UpdateSetmealRequest) (resp *setmeal.UpdateSetmealResponse, apiError *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service SetmealServiceImpl) GetSetmealPage(ctx context.Context, req *setmeal.GetSetmealsPageRequest) (resp *setmeal.GetSetmealsPageResponse, apiError *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service SetmealServiceImpl) ChangeSetmealStatus(ctx context.Context, req *setmeal.UpdateSetmealStatusRequest) (resp *setmeal.UpdateSetmealStatusResponse, apiError *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service SetmealServiceImpl) DeleteSetmeals(ctx context.Context, req *setmeal.DeleteSetmealsRequest) (resp *setmeal.DeleteSetmealsResponse, apiError *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service SetmealServiceImpl) CreateSetmeals(ctx context.Context, req *setmeal.AddSetmealRequest) (resp *setmeal.AddSetmealResponse, apiError *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func (service SetmealServiceImpl) GetSetmealsByID(ctx context.Context, req *setmeal.GetSetmealByIDRequest) (resp *setmeal.GetSetmealByIDResponse, apiError *model.ApiError) {
	//TODO implement me
	panic("implement me")
}

func NewSetmealServiceImpl(setmealDAO setmealDAO.SetmealDAOInterface) SetmealServiceImpl {
	return SetmealServiceImpl{setmealDAO}
}
