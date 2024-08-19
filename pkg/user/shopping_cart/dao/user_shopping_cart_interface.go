package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type UserShoppingCartDatabaseInterface interface {
	// AddToCart 将商品添加到数据库中的购物车表
	AddToCart(ctx context.Context, commodity *model.ShoppingCart) error

	// GetCartList 从数据库获取购物车列表
	GetCartList(ctx context.Context, userID uint) ([]model.ShoppingCart, error)

	// UpdateCart 更新数据库中的购物车项
	UpdateCart(ctx context.Context, commodity *model.ShoppingCart) error

	// DeleteCart 从数据库中删除某一项商品
	DeleteCart(ctx context.Context, userID uint, commodityID uint) error

	// DeleteCartByUserID 从数据库中删除某个人的所有购物车项
	DeleteCartByUserID(ctx context.Context, userID uint) error

	// GetCartTotalAmount 获取某个人的购物车中的所有商品的总价值
	GetCartTotalAmount(ctx context.Context, userID uint) (totalAmount float64, err error)

	// PersistCart 从 Redis 中获取购物车并保存到数据库
	PersistCart(ctx context.Context, userID uint) error
}

type UserShoppingCartDaoCacheInterface interface {
	// AddToCart 添加商品到购物车
	AddToCart(ctx context.Context, commodity *model.ShoppingCart) error

	// GetCartList 获取购物车列表
	GetCartList(ctx context.Context, userID uint) ([]model.ShoppingCart, error)

	// DeleteCart 删除购物车中的某一项商品
	DeleteCart(ctx context.Context, userID uint, commodity *model.ShoppingCart) error

	// DeleteCartByUserID 删除某个人的购物车中的所有一项商品
	DeleteCartByUserID(ctx context.Context, userID uint) error

	// GetCartTotalAmount 获取某个人的购物车中的所有商品的总价值
	GetCartTotalAmount(ctx context.Context, userID uint) (totalAmount float64, err error)
}
