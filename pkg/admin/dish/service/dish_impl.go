package service

import (
	"context"
	"encoding/json"
	"sky-take-out-gin/internal/utils/convert"
	model "sky-take-out-gin/model/sql"
	paramModel "sky-take-out-gin/pkg/admin/dish/DTO"
	dishDao "sky-take-out-gin/pkg/admin/dish/dao"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
)

type DishServiceImpl struct {
	dao dishDao.DishDaoInterface
}

func (service DishServiceImpl) UpdateDish(ctx context.Context, req *paramModel.UpdateDishRequest) (*paramModel.UpdateDishResponse, *error2.ApiError) {
	// 查找菜品并预加载口味
	dish, err := service.dao.SearchDishByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SearchDishByIDError,
			Msg:  err.Error(),
		}
	}

	// 更新菜品字段
	err = convert.UpdateStructFields(req, dish)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.UpdateDishError,
			Msg:  err.Error(),
		}
	}

	dish.UpdateUser = ctx.Value("userID").(uint)

	// 开始事务
	tx := service.dao.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// 更新菜品
	err = service.dao.UpdateDishWithTransaction(ctx, tx, dish)
	if err != nil {
		tx.Rollback()
		return nil, &error2.ApiError{
			Code: code.UpdateDishError,
			Msg:  err.Error(),
		}
	}

	// 删除旧的口味数据
	err = service.dao.DeleteDishFlavorsByDishIDWithTransaction(ctx, tx, dish.ID)
	if err != nil {
		tx.Rollback()
		return nil, &error2.ApiError{
			Code: code.DeleteDishFlavorError,
			Msg:  err.Error(),
		}
	}

	// 插入新的口味数据
	for _, flavorDTO := range req.Flavors {
		flavorValue, err := json.Marshal(flavorDTO.Value)
		if err != nil {
			tx.Rollback()
			return nil, &error2.ApiError{
				Code: code.CreateDishFlavorError,
				Msg:  err.Error(),
			}
		}

		flavor := model.DishFlavor{
			DishID: dish.ID,
			Name:   flavorDTO.Name,
			Value:  json.RawMessage(flavorValue),
		}
		err = service.dao.CreateDishFlavorWithTransaction(ctx, tx, flavor)
		if err != nil {
			tx.Rollback()
			return nil, &error2.ApiError{
				Code: code.CreateDishFlavorError,
				Msg:  err.Error(),
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &error2.ApiError{
			Code: code.UpdateDishError,
			Msg:  err.Error(),
		}
	}

	return &paramModel.UpdateDishResponse{}, nil
}
func (service DishServiceImpl) DeleteDish(ctx context.Context, req *paramModel.DeleteDishRequest) (resp *paramModel.DeleteDishResponse, apiError *error2.ApiError) {
	err := service.dao.DeleteDish(ctx, req.IDs)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.DeleteDishError,
			Msg:  err.Error(),
		}
	}
	return &paramModel.DeleteDishResponse{}, nil
}

func (service DishServiceImpl) AddDish(ctx context.Context, req *paramModel.AddDishRequest) (*paramModel.AddDishResponse, *error2.ApiError) {
	// 开始事务
	tx := service.dao.BeginTransaction()

	// 创建 Dish 实例
	dish := &model.Dish{}
	err := convert.UpdateStructFields(req, dish)
	if err != nil {
		tx.Rollback()
		return nil, &error2.ApiError{
			Code: code.CreateDishError,
			Msg:  err.Error(),
		}
	}

	// 设置创建人
	dish.CreateUser = ctx.Value("userID").(uint)

	// 插入 Dish 数据
	err = service.dao.CreateDishWithTransaction(ctx, tx, dish)
	if err != nil {
		tx.Rollback()
		return nil, &error2.ApiError{
			Code: code.CreateDishError,
			Msg:  err.Error(),
		}
	}

	// 插入 DishFlavor 数据
	for _, flavorDTO := range req.Flavors {
		// 将 Flavor 值转换为 JSON
		flavorValue, err := json.Marshal(flavorDTO.Value)
		if err != nil {
			tx.Rollback()
			return nil, &error2.ApiError{
				Code: code.CreateDishFlavorError,
				Msg:  err.Error(),
			}
		}

		// 创建 DishFlavor 实例
		flavor := model.DishFlavor{
			DishID: dish.ID, // 使用插入成功后的 DishID
			Name:   flavorDTO.Name,
			Value:  json.RawMessage(flavorValue),
		}

		// 插入 DishFlavor 数据
		err = service.dao.CreateDishFlavorWithTransaction(ctx, tx, flavor)
		if err != nil {
			tx.Rollback()
			return nil, &error2.ApiError{
				Code: code.CreateDishFlavorError,
				Msg:  err.Error(),
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &error2.ApiError{
			Code: code.CreateDishError,
			Msg:  err.Error(),
		}
	}

	return &paramModel.AddDishResponse{}, nil
}

func (service DishServiceImpl) SearchDishByID(ctx context.Context, req *paramModel.SearchDishByIDRequest) (resp *paramModel.SearchDishByIDResponse, apiError *error2.ApiError) {
	dish, err := service.dao.SearchDishByID(ctx, req.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SearchDishByIDError,
			Msg:  err.Error(),
		}
	}
	return &paramModel.SearchDishByIDResponse{Dish: *dish}, nil
}

func (service DishServiceImpl) SearchDishByCategory(ctx context.Context, req *paramModel.SearchDishByCategoryRequest) (resp *paramModel.SearchDishByCategoryResponse, apiError *error2.ApiError) {
	dish, err := service.dao.SearchDishByCategory(ctx, req.CategoryID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SearchDishByCategoryError,
			Msg:  err.Error(),
		}
	}
	return &paramModel.SearchDishByCategoryResponse{Records: dish}, nil
}

func (service DishServiceImpl) SearchDishByPage(ctx context.Context, req *paramModel.SearchDishByPageRequest) (resp *paramModel.SearchDishByPageResponse, apiError *error2.ApiError) {
	total, dishes, err := service.dao.SearchDishByPage(ctx, req.CategoryID, req.Name, req.Status, req.Page, req.PageSize)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SearchDishByPageError,
			Msg:  err.Error(),
		}
	}
	return &paramModel.SearchDishByPageResponse{Total: total, Records: dishes}, nil
}

func (service DishServiceImpl) ChangeDishStatus(ctx context.Context, req *paramModel.ChangeDishStatusRequest) (resp *paramModel.ChangeDishStatusResponse, apiError *error2.ApiError) {
	err := service.dao.ChangeDishStatus(ctx, req.ID, req.Status)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.ChangeDishStatusError,
			Msg:  err.Error(),
		}
	}
	return &paramModel.ChangeDishStatusResponse{}, nil
}

func NewDishServiceImpl(dao dishDao.DishDaoInterface) DishServiceImpl {
	return DishServiceImpl{dao: dao}
}
