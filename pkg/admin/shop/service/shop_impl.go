package service

import (
	"context"
	"sky-take-out-gin/pkg/admin/shop/DTO"
	shopDao "sky-take-out-gin/pkg/admin/shop/dao"
	error2 "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
)

type ShopServiceImpl struct {
	dao shopDao.ShopDaoInterface
}

func (service ShopServiceImpl) GetShopStatus(ctx context.Context, req *DTO.GetShopStatusRequest) (*DTO.GetShopStatusResponse, *error2.ApiError) {
	status, err := service.dao.GetShopStatus(ctx)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.GetShopStatusError,
			Msg:  err.Error(),
		}
	}
	return &DTO.GetShopStatusResponse{
		Status: status,
	}, nil
}

func (service ShopServiceImpl) SetShopStatus(ctx context.Context, req *DTO.SetShopStatusRequest) (*DTO.SetShopStatusResponse, *error2.ApiError) {
	err := service.dao.SetShopStatus(ctx, req.Status)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SetShopStatusError,
			Msg:  err.Error(),
		}
	}
	return &DTO.SetShopStatusResponse{}, nil
}

func NewShopServiceImpl(dao shopDao.ShopDaoInterface) ShopServiceImpl {
	return ShopServiceImpl{dao: dao}
}
