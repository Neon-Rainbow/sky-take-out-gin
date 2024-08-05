package DTO

type GetShopStatusRequest struct {
}

type GetShopStatusResponse struct {
	Status int `json:"status"`
}

type SetShopStatusRequest struct {
	Status int `uri:"status" json:"status"`
}

type SetShopStatusResponse struct {
}
