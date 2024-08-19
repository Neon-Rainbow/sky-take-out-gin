package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type UserShoppingCartRedisImpl struct {
	db database.DatabaseInterface
}

func (dao *UserShoppingCartRedisImpl) AddToCart(ctx context.Context, commodity *model.ShoppingCart) error {
	key := fmt.Sprintf("cart:%d", commodity.UserID)
	field := dao.getFieldKey(commodity)

	// 从 Redis 中获取当前的商品信息（JSON 字符串）
	result := dao.db.GetRedis().HGet(ctx, key, field)
	if err := result.Err(); err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("failed to get current commodity: %v", err)
	}

	var existingCommodity model.ShoppingCart
	if result.Val() != "" {
		if err := json.Unmarshal([]byte(result.Val()), &existingCommodity); err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("failed to unmarshal commodity: %v", err)
		}
	}

	// 更新数量
	commodity.Number += existingCommodity.Number

	// 将更新后的商品数量存回 Redis
	commodityJSON, err := json.Marshal(commodity)
	if err != nil {
		return fmt.Errorf("failed to marshal commodity: %v", err)
	}
	err = dao.db.GetRedis().HSet(ctx, key, field, commodityJSON).Err()
	return err
}

// GetCartList 从购物车中获取所有商品
func (dao *UserShoppingCartRedisImpl) GetCartList(ctx context.Context, userID uint) ([]model.ShoppingCart, error) {
	key := fmt.Sprintf("cart:%d", userID)
	items, err := dao.db.GetRedis().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get cart list: %v", err)
	}

	var cartList []model.ShoppingCart
	for _, itemJSON := range items {
		var item model.ShoppingCart
		if err := json.Unmarshal([]byte(itemJSON), &item); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cart item: %v", err)
		}
		cartList = append(cartList, item)
	}

	return cartList, nil
}

// DeleteCart 从购物车中删除商品，减少商品数量
func (dao *UserShoppingCartRedisImpl) DeleteCart(ctx context.Context, userID uint, commodity *model.ShoppingCart) error {
	key := fmt.Sprintf("cart:%d", userID)
	field := dao.getFieldKey(commodity)
	var err error

	result := dao.db.GetRedis().HGet(ctx, key, field)
	if err := result.Err(); err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("failed to get current commodity: %v", err)
	}

	var existingCommodity model.ShoppingCart
	if result.Val() != "" {
		if err := json.Unmarshal([]byte(result.Val()), &existingCommodity); err != nil && !errors.Is(err, redis.Nil) {
			return fmt.Errorf("failed to unmarshal commodity: %v", err)
		}
	}
	currentQuantity := existingCommodity.Number

	// 如果数量为零或者减去数量后为零，则从购物车中删除该商品
	if currentQuantity <= commodity.Number {
		err = dao.db.GetRedis().HDel(ctx, key, field).Err()
	} else {
		// 否则减少商品数量
		commodity.Number = currentQuantity - commodity.Number
		commodityJSON, err := json.Marshal(commodity)
		if err != nil {
			return fmt.Errorf("failed to marshal commodity: %v", err)
		}
		err = dao.db.GetRedis().HSet(ctx, key, field, commodityJSON).Err()
	}
	return err
}

// DeleteCartByUserID 删除用户的整个购物车
func (dao *UserShoppingCartRedisImpl) DeleteCartByUserID(ctx context.Context, userID uint) error {
	key := fmt.Sprintf("cart:%d", userID)
	err := dao.db.GetRedis().Del(ctx, key).Err()
	return err
}

// GetCartTotalAmount 获取购物车总金额
func (dao *UserShoppingCartRedisImpl) GetCartTotalAmount(ctx context.Context, userID uint) (totalAmount float64, err error) {
	cartList, err := dao.GetCartList(ctx, userID)
	if err != nil {
		return 0, err
	}

	for _, item := range cartList {
		totalAmount += item.Amount * float64(item.Number)
	}

	return totalAmount, nil
}

// 辅助方法，根据商品类型获取在Redis中的键
func (dao *UserShoppingCartRedisImpl) getFieldKey(commodity *model.ShoppingCart) string {
	if commodity.DishID != 0 {
		return fmt.Sprintf("dish:%d-flavor:%s", commodity.DishID, commodity.DishFlavor)
	} else if commodity.SetmealID != 0 {
		return fmt.Sprintf("setmeal:%d", commodity.SetmealID)
	}
	return ""
}

func NewUserShoppingCartRedisImpl(db database.DatabaseInterface) *UserShoppingCartRedisImpl {
	return &UserShoppingCartRedisImpl{db: db}
}
