package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	error2 "sky-take-out-gin/pkg/common/error"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
	paramModel "sky-take-out-gin/pkg/shop/DTO"
	shopService "sky-take-out-gin/pkg/shop/service"
)

type ShopControllerImpl struct {
	shopService shopService.ShopServiceInterface
}

func (s ShopControllerImpl) GetShopStatus(c *gin.Context) {
	req := paramModel.GetShopStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return s.shopService.GetShopStatus(ctx, req.(*paramModel.GetShopStatusRequest))
		},
	)
}

func (s ShopControllerImpl) SetShopStatus(c *gin.Context) {
	req := paramModel.SetShopStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *error2.ApiError) {
			return s.shopService.SetShopStatus(ctx, req.(*paramModel.SetShopStatusRequest))
		},
		c.ShouldBindUri,
	)
}

func NewShopControllerImpl(shopService shopService.ShopServiceInterface) ShopControllerImpl {
	return ShopControllerImpl{shopService: shopService}
}
