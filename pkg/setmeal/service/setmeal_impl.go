package service

import (
	"context"
	"sky-take-out-gin/internal/utils/convert"
	sqlModel "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/setmeal/DTO"
	setmealDAO "sky-take-out-gin/pkg/setmeal/dao"
)

type SetmealServiceImpl struct {
	dao setmealDAO.SetmealDAOInterface
}

func (service SetmealServiceImpl) UpdateSetmeal(ctx context.Context, req *DTO.UpdateSetmealRequest) (resp *DTO.UpdateSetmealResponse, apiError *error2.ApiError) {
	s := &sqlModel.Setmeal{}
	_ = convert.UpdateStructFields(req, s)

	s.UpdateUser = ctx.Value("userID").(uint)

	err := service.dao.UpdateSetmeal(ctx, s)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &DTO.UpdateSetmealResponse{}, nil
}

func (service SetmealServiceImpl) GetSetmealPage(ctx context.Context, req *DTO.GetSetmealsPageRequest) (resp *DTO.GetSetmealsPageResponse, apiError *error2.ApiError) {
	records, err := service.dao.GetSetmealPage(ctx, req.CategoryID, req.Page, req.PageSize)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.GetSetmealPageError,
			Msg:  err.Error(),
		}
	}
	return &DTO.GetSetmealsPageResponse{
		Total:   len(records),
		Records: records,
	}, nil
}

func (service SetmealServiceImpl) ChangeSetmealStatus(ctx context.Context, req *DTO.UpdateSetmealStatusRequest) (resp *DTO.UpdateSetmealStatusResponse, apiError *error2.ApiError) {
	err := service.dao.UpdateSetmealStatus(ctx, req.ID, req.Status)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &DTO.UpdateSetmealStatusResponse{}, nil
}

func (service SetmealServiceImpl) DeleteSetmeals(ctx context.Context, req *DTO.DeleteSetmealsRequest) (resp *DTO.DeleteSetmealsResponse, apiError *error2.ApiError) {
	err := service.dao.DeleteSetmeals(ctx, req.IDs)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &DTO.DeleteSetmealsResponse{}, nil
}

func (service SetmealServiceImpl) CreateSetmeals(ctx context.Context, req *DTO.AddSetmealRequest) (resp *DTO.AddSetmealResponse, apiError *error2.ApiError) {
	var s sqlModel.Setmeal
	err := convert.UpdateStructFields(req, &s)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.ParamError,
			Msg:  err.Error(),
		}
	}

	s.CreateUser = ctx.Value("userID").(uint)
	s.UpdateUser = ctx.Value("userID").(uint)
	err = service.dao.CreateSetmeal(ctx, &s)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.UpdateSetmealError,
			Msg:  err.Error(),
		}
	}
	return &DTO.AddSetmealResponse{}, nil
}

func (service SetmealServiceImpl) GetSetmealsByID(ctx context.Context, req *DTO.GetSetmealByIDRequest) (resp *DTO.GetSetmealByIDResponse, apiError *error2.ApiError) {
	s, err := service.dao.GetSetmealByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.GetSetmealPageError,
			Msg:  err.Error(),
		}
	}
	resp = &DTO.GetSetmealByIDResponse{}

	err = convert.UpdateStructFields(s, resp)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.ParamError,
			Msg:  err.Error(),
		}
	}
	return resp, nil
}

func NewSetmealServiceImpl(setmealDAO setmealDAO.SetmealDAOInterface) SetmealServiceImpl {
	return SetmealServiceImpl{setmealDAO}
}
