package DTO

// QueryParams 查询参数
// 用于条件查询订单
type QueryParams struct {
	BeginTime string `json:"begin_time" form:"begin_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	OrderID   int    `json:"order_id" form:"order_id"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	Phone     string `json:"phone" form:"phone"`
	Status    int    `json:"status" form:"status"`
}

// NewQueryParams 创建一个默认 QueryParams 实例
func NewQueryParams() *QueryParams {
	return &QueryParams{
		Page:     1,  // 默认值为 1
		PageSize: 10, // 默认值为 10
	}
}

type CancelOrderRequestDTO struct {
	OrderID      uint   `json:"order_id" binding:"required"`
	CancelReason string `json:"cancel_reason" binding:"required"`
}

type RejectOrderRequestDTO struct {
	OrderID      uint   `json:"order_id" binding:"required"`
	RejectReason string `json:"reject_reason" binding:"required"`
}

type OrderIDDTO struct {
	OrderID uint `json:"order_id"  uri:"order_id" form:"order_id" binding:"required"`
}
