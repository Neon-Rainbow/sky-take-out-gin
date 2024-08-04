package shop

import (
	"context"
	"sky-take-out-gin/code"
	shopDao "sky-take-out-gin/internal/dao/admin/shop"
	"sky-take-out-gin/model"
	"sky-take-out-gin/model/param/admin/shop"
)

type ShopServiceImpl struct {
	dao shopDao.ShopDaoInterface
}

func (service ShopServiceImpl) GetShopStatus(ctx context.Context, req *shop.GetShopStatusRequest) (*shop.GetShopStatusResponse, *model.ApiError) {
	status, err := service.dao.GetShopStatus(ctx)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.GetShopStatusError,
			Msg:  err.Error(),
		}
	}
	return &shop.GetShopStatusResponse{
		Status: status,
	}, nil
}

func (service ShopServiceImpl) SetShopStatus(ctx context.Context, req *shop.SetShopStatusRequest) (*shop.SetShopStatusResponse, *model.ApiError) {
	err := service.dao.SetShopStatus(ctx, req.Status)
	if err != nil {
		return nil, &model.ApiError{
			Code: code.SetShopStatusError,
			Msg:  err.Error(),
		}
	}
	return &shop.SetShopStatusResponse{}, nil
}

func NewShopServiceImpl(dao shopDao.ShopDaoInterface) ShopServiceImpl {
	return ShopServiceImpl{dao: dao}
}
