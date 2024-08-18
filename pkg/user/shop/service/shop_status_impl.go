package service

import (
	"context"
	apiErrorModel "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/user/shop/dao"
)

type ShopStatusServiceImpl struct {
	dao dao.ShopStatusDaoInterface
}

// GetShopStatus 获取店铺状态
func (service *ShopStatusServiceImpl) GetShopStatus(ctx context.Context) (status int, apiError *apiErrorModel.ApiError) {
	var err error
	status, err = service.dao.GetShopStatus(ctx)
	if err != nil {
		return -1, &apiErrorModel.ApiError{
			Code: code.GetShopStatusError,
			Msg:  err.Error(),
		}
	}
	return status, nil
}

// NewShopStatusServiceImpl 创建一个新的 ShopStatusServiceImpl
func NewShopStatusServiceImpl(dao dao.ShopStatusDaoInterface) *ShopStatusServiceImpl {
	return &ShopStatusServiceImpl{dao: dao}
}
