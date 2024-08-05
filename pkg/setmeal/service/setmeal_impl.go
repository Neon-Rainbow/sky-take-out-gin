package setmeal

import (
	"context"
	"sky-take-out-gin/code"
	setmealDAO "sky-take-out-gin/internal/dao/admin/setmeal"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/setmeal"
	sqlModel "sky-take-out-gin/model/sql"
	"sky-take-out-gin/utils/convert"
	"time"
)

type SetmealServiceImpl struct {
	dao setmealDAO.SetmealDAOInterface
}

func (service SetmealServiceImpl) UpdateSetmeal(ctx context.Context, req *setmeal.UpdateSetmealRequest) (resp *setmeal.UpdateSetmealResponse, apiError *model.ApiError) {
	s := &sqlModel.Setmeal{}
	_ = convert.UpdateStructFields(req, s)

	s.UpdateUser = ctx.Value("userID").(int64)
	s.UpdateTime = time.Now()

	err := service.dao.UpdateSetmeal(ctx, s)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &setmeal.UpdateSetmealResponse{}, nil
}

func (service SetmealServiceImpl) GetSetmealPage(ctx context.Context, req *setmeal.GetSetmealsPageRequest) (resp *setmeal.GetSetmealsPageResponse, apiError *model.ApiError) {
	records, err := service.dao.GetSetmealPage(ctx, req.CategoryID, req.Page, req.PageSize)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.GetSetmealPageError,
			Msg:  err.Error(),
		}
	}
	return &setmeal.GetSetmealsPageResponse{
		Total:   len(records),
		Records: records,
	}, nil
}

func (service SetmealServiceImpl) ChangeSetmealStatus(ctx context.Context, req *setmeal.UpdateSetmealStatusRequest) (resp *setmeal.UpdateSetmealStatusResponse, apiError *model.ApiError) {
	err := service.dao.UpdateSetmealStatus(ctx, req.ID, req.Status)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &setmeal.UpdateSetmealStatusResponse{}, nil
}

func (service SetmealServiceImpl) DeleteSetmeals(ctx context.Context, req *setmeal.DeleteSetmealsRequest) (resp *setmeal.DeleteSetmealsResponse, apiError *model.ApiError) {
	err := service.dao.DeleteSetmeals(ctx, req.IDs)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &setmeal.DeleteSetmealsResponse{}, nil
}

func (service SetmealServiceImpl) CreateSetmeals(ctx context.Context, req *setmeal.AddSetmealRequest) (resp *setmeal.AddSetmealResponse, apiError *model.ApiError) {
	var s sqlModel.Setmeal
	err := convert.UpdateStructFields(req, &s)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.ParamError,
			Msg:  err.Error(),
		}
	}

	s.CreateTime = time.Now()
	s.UpdateTime = time.Now()
	s.CreateUser = ctx.Value("userID").(int64)
	s.UpdateUser = ctx.Value("userID").(int64)
	err = service.dao.CreateSetmeal(ctx, &s)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &setmeal.AddSetmealResponse{}, nil
}

func (service SetmealServiceImpl) GetSetmealsByID(ctx context.Context, req *setmeal.GetSetmealByIDRequest) (resp *setmeal.GetSetmealByIDResponse, apiError *model.ApiError) {
	s, err := service.dao.GetSetmealByID(ctx, req.ID)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.GetSetmealPageError,
			Msg:  err.Error(),
		}
	}
	resp = &setmeal.GetSetmealByIDResponse{}

	err = convert.UpdateStructFields(s, resp)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.ParamError,
			Msg:  err.Error(),
		}
	}
	return resp, nil
}

func NewSetmealServiceImpl(setmealDAO setmealDAO.SetmealDAOInterface) SetmealServiceImpl {
	return SetmealServiceImpl{setmealDAO}
}
