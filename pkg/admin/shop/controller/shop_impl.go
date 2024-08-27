package controller

import (
	"github.com/gin-gonic/gin"
	shopService "sky-take-out-gin/pkg/admin/shop/service"
	HandleRequest "sky-take-out-gin/pkg/common/request_handle"
)

type ShopControllerImpl struct {
	shopService shopService.ShopServiceInterface
}

func (s ShopControllerImpl) GetShopStatus(c *gin.Context) {

	HandleRequest.HandleRequest(c, s.shopService.GetShopStatus)
}

func (s ShopControllerImpl) SetShopStatus(c *gin.Context) {

	HandleRequest.HandleRequest(c, s.shopService.SetShopStatus, c.ShouldBindUri)
}

func NewShopControllerImpl(shopService shopService.ShopServiceInterface) ShopControllerImpl {
	return ShopControllerImpl{shopService: shopService}
}
