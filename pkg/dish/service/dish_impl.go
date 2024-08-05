package service

import (
	"context"
	"sky-take-out-gin/internal/utils/convert"
	"sky-take-out-gin/pkg/common/code"
	error2 "sky-take-out-gin/pkg/common/error"
	paramModel "sky-take-out-gin/pkg/dish/DTO"
	dishDao "sky-take-out-gin/pkg/dish/dao"
	"time"
)

type DishServiceImpl struct {
	dao dishDao.DishDaoInterface
}

func (service DishServiceImpl) UpdateDish(ctx context.Context, req *paramModel.UpdateDishRequest) (resp *paramModel.UpdateDishResponse, apiError *error2.ApiError) {
	dish, err := service.dao.SearchDishByID(ctx, req.Dish.ID)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.SearchDishByIDError,
			Msg:  err.Error(),
		}
	}

	err = convert.UpdateStructFields(req.Dish, dish)
	if err != nil {
		return nil, &error2.ApiError{
			Code: code.UpdateDishError,
			Msg:  err.Error(),
		}
	}

	dish.UpdateUser = ctx.Value("userID").(int64)
	dish.UpdateTime = time.Now()

	err = service.dao.UpdateDish(ctx, *dish)
	if err != nil {
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

func (service DishServiceImpl) AddDish(ctx context.Context, req *paramModel.AddDishRequest) (resp *paramModel.AddDishResponse, apiError *error2.ApiError) {
	err := service.dao.CreateDish(ctx, req.Dish)
	if err != nil {
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
