package shop

import (
	"context"
	"github.com/gin-gonic/gin"
	HandleRequest "sky-take-out-gin/internal/controller"
	shopService "sky-take-out-gin/internal/service/admin/shop"
	"sky-take-out-gin/model"
	paramModel "sky-take-out-gin/model/param/admin/shop"
)

type ShopControllerImpl struct {
	shopService shopService.ShopServiceInterface
}

func (s ShopControllerImpl) GetShopStatus(c *gin.Context) {
	req := paramModel.GetShopStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return s.shopService.GetShopStatus(ctx, req.(*paramModel.GetShopStatusRequest))
		},
	)
}

func (s ShopControllerImpl) SetShopStatus(c *gin.Context) {
	req := paramModel.SetShopStatusRequest{}
	HandleRequest.HandleRequest(
		c,
		&req,
		func(ctx context.Context, req interface{}) (interface{}, *model.ApiError) {
			return s.shopService.SetShopStatus(ctx, req.(*paramModel.SetShopStatusRequest))
		},
		c.ShouldBindUri,
	)
}

func NewShopControllerImpl(shopService shopService.ShopServiceInterface) ShopControllerImpl {
	return ShopControllerImpl{shopService: shopService}
}
