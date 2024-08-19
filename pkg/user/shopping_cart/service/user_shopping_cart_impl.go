package service

import (
	"context"
	"fmt"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/common/cache"
	"sky-take-out-gin/pkg/common/code"
	dao2 "sky-take-out-gin/pkg/user/dish/dao"
	dao3 "sky-take-out-gin/pkg/user/set_meal/dao"
	"sky-take-out-gin/pkg/user/shopping_cart/DTO"
	"sky-take-out-gin/pkg/user/shopping_cart/dao"
	"time"
)

type UserShoppingCartServiceImpl struct {
	daoCache    dao.UserShoppingCartDaoCacheInterface
	daoDatabase dao.UserShoppingCartDatabaseInterface
	cache       cache.RedisCacheInterface
	daoDish     dao2.DishDaoInterface
	daoSetMeal  dao3.SetMealDaoInterface
}

// AddToCart 添加商品到购物车
func (service *UserShoppingCartServiceImpl) AddToCart(ctx context.Context, userID uint, req *DTO.ShoppingCartRequest) *api_error.ApiError {
	// 构建 ShoppingCart 对象

	dish := model.Dish{}
	setMeal := model.Setmeal{}

	if req.DishID != 0 {
		err := service.cache.GetOrSet(
			ctx,
			fmt.Sprintf("dish:%d", req.DishID),
			time.Hour,
			&dish,
			func(ctx context.Context, args ...interface{}) (interface{}, error) {
				return service.daoDish.GetDishByDishID(ctx, req.DishID)
			},
		)
		if err != nil {
			return api_error.NewApiError(code.GetDishByIdError, fmt.Sprintf("failed to get dish: %v", err))
		}
	} else {
		err := service.cache.GetOrSet(
			ctx,
			fmt.Sprintf("set_meal:%d", req.SetMealID),
			time.Hour,
			&setMeal,
			func(ctx context.Context, args ...interface{}) (interface{}, error) {
				return service.daoSetMeal.GetSetMealBySetMealID(ctx, req.SetMealID)
			},
		)
		if err != nil {
			return api_error.NewApiError(code.SetMealGetDetailError, fmt.Sprintf("failed to get set meal: %v", err))
		}
	}
	var price float64
	if req.DishID != 0 {
		price = dish.Price
	} else {
		price = setMeal.Price
	}
	var name string
	if req.DishID != 0 {
		name = dish.Name
	} else {
		name = setMeal.Name
	}

	var image string
	if req.DishID != 0 {
		image = dish.Image
	} else {
		image = setMeal.Image
	}

	commodity := &model.ShoppingCart{
		UserID:     userID,
		DishID:     req.DishID,
		DishFlavor: req.DishFlavor,
		SetmealID:  req.SetMealID,
		Number:     1,
		Amount:     price,
		Name:       name,
		Image:      image,
	}

	// 更新 Redis 缓存
	if err := service.daoCache.AddToCart(ctx, commodity); err != nil {
		return api_error.NewApiError(code.AddCartToCacheError, fmt.Sprintf("failed to add to cart in cache: %v", err))
	}

	// 可选择在此更新数据库
	// err := service.daoDatabase.AddToCart(ctx, commodity)
	// if err != nil {
	//     return api_error.NewInternalError(fmt.Sprintf("failed to add to cart in database: %v", err))
	// }

	return nil
}

// GetCartList 获取购物车列表
func (service *UserShoppingCartServiceImpl) GetCartList(ctx context.Context, userID uint) ([]model.ShoppingCart, *api_error.ApiError) {
	// 从 Redis 缓存中获取购物车列表
	cartList, err := service.daoCache.GetCartList(ctx, userID)
	if err != nil {
		return nil, api_error.NewApiError(code.GetCartListFromCacheError, fmt.Sprintf("failed to get cart list from cache: %v", err))
	}

	// 如果缓存中没有数据，可以选择从数据库中获取并更新缓存
	// if len(cartList) == 0 {
	//     cartList, err = service.daoDatabase.GetCartList(ctx, userID)
	//     if err != nil {
	//         return nil, api_error.NewInternalError(fmt.Sprintf("failed to get cart list from database: %v", err))
	//     }
	//     // 更新缓存
	//     // err = service.daoCache.SetCartList(ctx, userID, cartList)
	// }

	return cartList, nil
}

// DeleteCart 删除购物车中的指定商品
func (service *UserShoppingCartServiceImpl) DeleteCart(ctx context.Context, userID uint, req *DTO.ShoppingCartRequest) *api_error.ApiError {
	commodity := &model.ShoppingCart{
		UserID:     userID,
		DishID:     req.DishID,
		DishFlavor: req.DishFlavor,
		Number:     1,
		SetmealID:  req.SetMealID,
	}

	// 更新 Redis 缓存
	if err := service.daoCache.DeleteCart(ctx, userID, commodity); err != nil {
		return api_error.NewApiError(code.DeleteCartFromCacheError, fmt.Sprintf("failed to delete from cart in cache: %v", err))
	}

	// 可选择在此更新数据库
	// err := service.daoDatabase.DeleteCart(ctx, userID, commodity)
	// if err != nil {
	//     return api_error.NewInternalError(fmt.Sprintf("failed to delete from cart in database: %v", err))
	// }

	return nil
}

// DeleteCartByUserID 删除用户的整个购物车
func (service *UserShoppingCartServiceImpl) DeleteCartByUserID(ctx context.Context, userID uint) *api_error.ApiError {
	// 更新 Redis 缓存
	if err := service.daoCache.DeleteCartByUserID(ctx, userID); err != nil {
		return api_error.NewApiError(code.DeleteCartFromCacheError, fmt.Sprintf("failed to delete cart by user ID in cache: %v", err))
	}

	// 可选择在此更新数据库
	// err := service.daoDatabase.DeleteCartByUserID(ctx, userID)
	// if err != nil {
	//     return api_error.NewInternalError(fmt.Sprintf("failed to delete cart by user ID in database: %v", err))
	// }

	return nil
}

// GetCartTotalAmount 获取购物车总金额
func (service *UserShoppingCartServiceImpl) GetCartTotalAmount(ctx context.Context, userID uint) (totalAmount float64, apiErr *api_error.ApiError) {
	// 从 Redis 缓存中获取购物车总金额
	totalAmount, err := service.daoCache.GetCartTotalAmount(ctx, userID)
	if err != nil {
		return 0, api_error.NewApiError(code.GetCartTotalAmountFromRedisError, fmt.Sprintf("failed to get cart total amount from cache: %v", err))
	}

	// 可选择从数据库获取并更新缓存
	// if totalAmount == 0 {
	//     totalAmount, err = service.daoDatabase.GetCartTotalAmount(ctx, userID)
	//     if err != nil {
	//         return 0, api_error.NewInternalError(fmt.Sprintf("failed to get cart total amount from database: %v", err))
	//     }
	//     // 更新缓存
	//     // err = service.daoCache.SetCartTotalAmount(ctx, userID, totalAmount)
	// }

	return totalAmount, nil
}

// PersistCart 将购物车从缓存中持久化到数据库
func (service *UserShoppingCartServiceImpl) PersistCart(ctx context.Context, userID uint) *api_error.ApiError {
	// 从 Redis 缓存中获取购物车列表
	cartList, err := service.daoCache.GetCartList(ctx, userID)
	if err != nil {
		return api_error.NewApiError(code.SaveRedisToDatabaseError, fmt.Sprintf("failed to get cart list from cache: %v", err))
	}

	// 将购物车列表保存到数据库
	for _, item := range cartList {
		if err := service.daoDatabase.AddToCart(ctx, &item); err != nil {
			return api_error.NewApiError(code.SaveRedisToDatabaseError, fmt.Sprintf("failed to persist cart item to database: %v", err))
		}
	}

	return nil
}

func NewUserShoppingCartService(
	daoCache dao.UserShoppingCartDaoCacheInterface,
	daoDatabase dao.UserShoppingCartDatabaseInterface,
	cache cache.RedisCacheInterface,
	daoDish dao2.DishDaoInterface,
	daoSetMeal dao3.SetMealDaoInterface,
) *UserShoppingCartServiceImpl {
	return &UserShoppingCartServiceImpl{
		daoCache,
		daoDatabase,
		cache,
		daoDish,
		daoSetMeal,
	}
}
