package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/user/shopping_cart/DTO"
)

type UserShoppingCartServiceInterface interface {
	// AddToCart 添加商品到购物车
	AddToCart(ctx context.Context, userID uint, req *DTO.ShoppingCartRequest) *api_error.ApiError

	// GetCartList 获取购物车列表
	GetCartList(ctx context.Context, userID uint) ([]model.ShoppingCart, *api_error.ApiError)

	// DeleteCart 删除购物车中的某一项商品
	DeleteCart(ctx context.Context, userID uint, req *DTO.ShoppingCartRequest) *api_error.ApiError

	// DeleteCartByUserID 删除某个人的购物车中的所有一项商品
	DeleteCartByUserID(ctx context.Context, userID uint) *api_error.ApiError

	// GetCartTotalAmount 获取某个人的购物车中的所有商品的总价值
	GetCartTotalAmount(ctx context.Context, userID uint) (totalAmount float64, apiErr *api_error.ApiError)

	// PersistCart 从 Redis 中获取购物车并保存到数据库
	PersistCart(ctx context.Context, userID uint) *api_error.ApiError
}
