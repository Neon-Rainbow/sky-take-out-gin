package controller

import (
	"github.com/gin-gonic/gin"
	paramModel "sky-take-out-gin/pkg/admin/shop/DTO"
	shopService "sky-take-out-gin/pkg/admin/shop/service"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
)

type ShopControllerImpl struct {
	shopService shopService.ShopServiceInterface
}

func (s ShopControllerImpl) GetShopStatus(c *gin.Context) {
	req := paramModel.GetShopStatusRequest{}
	HandleRequest.HandleRequest(c, &req, s.shopService.GetShopStatus)
}

func (s ShopControllerImpl) SetShopStatus(c *gin.Context) {
	req := paramModel.SetShopStatusRequest{}
	HandleRequest.HandleRequest(c, &req, s.shopService.SetShopStatus, c.ShouldBindUri)
}

func NewShopControllerImpl(shopService shopService.ShopServiceInterface) ShopControllerImpl {
	return ShopControllerImpl{shopService: shopService}
}
