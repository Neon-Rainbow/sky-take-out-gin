package dao

import (
	"context"
	"encoding/json"
	"fmt"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type UserShoppingCartDatabaseImpl struct {
	db database.DatabaseInterface
}

func (dao *UserShoppingCartDatabaseImpl) AddToCart(ctx context.Context, commodity *model.ShoppingCart) error {
	err := dao.db.GetDB().WithContext(ctx).Create(commodity).Error
	return err
}

func (dao *UserShoppingCartDatabaseImpl) GetCartList(ctx context.Context, userID uint) ([]model.ShoppingCart, error) {
	var commodities []model.ShoppingCart
	err := dao.db.GetDB().WithContext(ctx).Where("user_id = ?", userID).Find(&commodities).Error
	return commodities, err
}

func (dao *UserShoppingCartDatabaseImpl) UpdateCart(ctx context.Context, commodity *model.ShoppingCart) error {
	err := dao.db.GetDB().WithContext(ctx).Save(commodity).Error
	return err
}

func (dao *UserShoppingCartDatabaseImpl) DeleteCart(ctx context.Context, userID uint, commodityID uint) error {
	err := dao.db.GetDB().WithContext(ctx).Where("id = ?", commodityID).Delete(&model.ShoppingCart{}).Error
	return err
}

func (dao *UserShoppingCartDatabaseImpl) DeleteCartByUserID(ctx context.Context, userID uint) error {
	err := dao.db.GetDB().WithContext(ctx).Where("user_id = ?", userID).Delete(&model.ShoppingCart{}).Error
	return err
}

func (dao *UserShoppingCartDatabaseImpl) GetCartTotalAmount(ctx context.Context, userID uint) (totalAmount float64, err error) {
	err = dao.db.GetDB().WithContext(ctx).Model(&model.ShoppingCart{}).Where("user_id = ?", userID).Select("sum(price * quantity)").Row().Scan(&totalAmount)
	return totalAmount, err
}

// PersistCart 从 Redis 中获取购物车并保存到数据库
func (dao *UserShoppingCartDatabaseImpl) PersistCart(ctx context.Context, userID uint) error {
	key := fmt.Sprintf("cart:%d", userID)
	redisData, err := dao.db.GetRedis().HGetAll(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get cart data from redis: %v", err)
	}

	if len(redisData) == 0 {
		return fmt.Errorf("no cart data found for user %d", userID)
	}

	// 2. 将购物车数据保存到数据库
	for _, itemJSON := range redisData {
		var item model.ShoppingCart
		if err := json.Unmarshal([]byte(itemJSON), &item); err != nil {
			return fmt.Errorf("failed to unmarshal cart item: %v", err)
		}
		item.UserID = userID // 确保UserID被正确设置
		if err := dao.AddToCart(ctx, &item); err != nil {
			return fmt.Errorf("failed to save cart item to database: %v", err)
		}
	}

	// 3. 清空 Redis 中的购物车
	if err := dao.db.GetRedis().Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to clear cart data in redis: %v", err)
	}
	return nil
}

func NewUserShoppingCartDatabaseImpl(db database.DatabaseInterface) *UserShoppingCartDatabaseImpl {
	return &UserShoppingCartDatabaseImpl{
		db: db,
	}
}
