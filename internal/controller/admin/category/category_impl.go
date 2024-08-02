package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/code"
	controllerResponse "sky-take-out-gin/internal/controller"
	serviceCategory "sky-take-out-gin/internal/service/admin/category"
	controllerModel "sky-take-out-gin/model"
	paramCategory "sky-take-out-gin/model/param/admin/category"
	"time"
)

type AdminCategoryControllerImpl struct {
	serviceCategory.CategoryService
}

func NewAdminCategoryControllerImpl(service serviceCategory.CategoryService) *AdminCategoryControllerImpl {
	return &AdminCategoryControllerImpl{
		service,
	}
}

// UpdateCategory 更新分类
// @Summary 更新分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param id body int true "分类ID"
// @Param name body string true "分类名称"
// @Param type body int true "分类类型"
// @Param sort body int true "分类排序"
// @Param status body int true "分类状态"
// @Success http.StatusOK {object} controller.Response
// @Failure http.StatusBadRequest {object} controller.Response
// @Failure http.StatusRequestTimeout {object} controller.Response
// @Failure http.StatusInternalServerError {object} controller.Response
// @Router /admin/category [put]
func (controller *AdminCategoryControllerImpl) UpdateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	resultChan := make(chan interface{})
	defer close(resultChan)

	go func() {

		// 从请求中获取参数
		var req paramCategory.AdminUpdateCategoryRequest
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryBindParamError,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryBindParamError,
					code.CategoryBindParamError.Message(),
					err.Error()),
			}
			return
		}

		// 转换参数
		category := req.ConvertToCategory()

		// 调用service层方法
		if err := controller.CategoryService.UpdateCategory(ctx, category); err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryUpdateFailed,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryUpdateFailed,
					code.CategoryUpdateFailed.Message(),
					err.Error()),
			}
			return
		}
		// 操作成功,返回响应
		resultChan <- &paramCategory.AdminUpdateCategoryResponse{}
		return
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			controllerResponse.ResponseErrorWithCode(c, http.StatusRequestTimeout, code.ServerError)
			return
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			controllerResponse.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
			return
		}
	case result := <-resultChan:
		switch res := result.(type) {
		case *controllerModel.ApiError:
			controllerResponse.ResponseErrorWithApiError(c, http.StatusBadRequest, res)
			return
		case *paramCategory.AdminUpdateCategoryResponse:
			controllerResponse.ResponseSuccess(c, res)
			return
		default:
			controllerResponse.ResponseErrorWithApiError(c, http.StatusInternalServerError, &controllerModel.ApiError{
				Code: code.ServerError,
				Msg:  fmt.Sprintf("未知类型错误, 类型为: %T", res),
			})
			return
		}
	}

}

// GetCategoryPage 分类分页查询
// @Summary 分类分页查询
// @Tags 分类
// @Accept json
// @Produce json
// @Param name query string false "分类名称"
// @Param page query int true "页码"
// @Param pageSize query int true "每页记录数"
// @Param type query int false "分类类型：1为菜品分类，2为套餐分类"
// @Success 200 {object} controller.Response
// @Failure 400 {object} controller.Response
// @Failure 404 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/category/page [get]
func (controller *AdminCategoryControllerImpl) GetCategoryPage(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	resultChan := make(chan interface{})
	defer close(resultChan)

	go func() {
		var req paramCategory.AdminCategoryPageRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryBindParamError,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryBindParamError,
					code.CategoryBindParamError.Message(),
					err.Error()),
			}
			return
		}

		// 调用service层方法
		res, err := controller.CategoryService.GetCategoryPage(ctx, &req)
		if err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryGetFailed,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryGetFailed,
					code.CategoryGetFailed.Message(),
					err.Error()),
			}
			return
		}
		resultChan <- res
		return
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			controllerResponse.ResponseErrorWithCode(c, http.StatusRequestTimeout, code.ServerError)
			return
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			controllerResponse.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
			return
		}
	case result := <-resultChan:
		switch res := result.(type) {
		case *controllerModel.ApiError:
			controllerResponse.ResponseErrorWithApiError(c, http.StatusBadRequest, res)
			return
		case *paramCategory.AdminCategoryPageResponse:
			controllerResponse.ResponseSuccess(c, res)
			return
		default:
			controllerResponse.ResponseErrorWithApiError(c, http.StatusInternalServerError, &controllerModel.ApiError{
				Code: code.ServerError,
				Msg:  fmt.Sprintf("未知类型错误, 类型为: %T", res),
			})
			return
		}
	}
}

// ChangeCategoryStatus 启用、禁用分类
// @Summary 启用、禁用分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param id query int true "分类ID"
// @Param status path int true "分类状态: 1为启用，0为禁用"
// @Success 200 {object} controller.Response
// @Failure 400 {object} controller.Response
// @Failure 404 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/category/status/{status} [post]
func (controller *AdminCategoryControllerImpl) ChangeCategoryStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	resultChan := make(chan interface{})
	defer close(resultChan)

	go func() {
		var req paramCategory.AdminChangeCategoryStatusRequest
		err := c.ShouldBind(&req)
		if err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryBindParamError,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryBindParamError,
					code.CategoryBindParamError.Message(),
					err.Error()),
			}
			return
		}

		// 调用service层方法
		err = controller.CategoryService.ChangeCategoryStatus(ctx, &req)
		if err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryChangeStatusFailed,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryChangeStatusFailed,
					code.CategoryChangeStatusFailed.Message(),
					err.Error()),
			}
			return
		}
		resultChan <- &paramCategory.AdminChangeCategoryStatusResponse{}
		return
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			controllerResponse.ResponseErrorWithCode(c, http.StatusRequestTimeout, code.ServerError)
			return
		}
		controllerResponse.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	case result := <-resultChan:
		switch res := result.(type) {
		case *controllerModel.ApiError:
			controllerResponse.ResponseErrorWithApiError(c, http.StatusBadRequest, res)
			return
		case *paramCategory.AdminChangeCategoryStatusResponse:
			controllerResponse.ResponseSuccess(c, res)
			return
		}

	}
}

// CreateCategory 新增分类
// @Summary 新增分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param id body int false "分类ID"
// @Param name body string true "分类名称"
// @Param type body int true "分类类型"
// @Param sort body int true "分类排序"
// @Success 200 {object} controller.Response
// @Failure 400 {object} controller.Response
// @Failure 404 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/category [post]
func (controller *AdminCategoryControllerImpl) CreateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	resultChan := make(chan interface{})
	defer close(resultChan)

	go func() {
		var req paramCategory.AdminCreateCategoryRequest
		err := c.ShouldBind(&req)
		if err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryBindParamError,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryBindParamError,
					code.CategoryBindParamError.Message(),
					err.Error()),
			}
			return
		}

		// 调用service层方法
		err = controller.CategoryService.CreateCategory(ctx, &req)
		if err != nil {
			resultChan <- &controllerModel.ApiError{
				Code: code.CategoryCreateFailed,
				Msg: fmt.Sprintf("Code: %d, Message: %s, Error detail: %s",
					code.CategoryCreateFailed,
					code.CategoryCreateFailed.Message(),
					err.Error()),
			}
			return
		}
		resultChan <- &paramCategory.AdminCreateCategoryResponse{}
		return
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			controllerResponse.ResponseErrorWithCode(c, http.StatusRequestTimeout, code.ServerError)
			return
		}
		controllerResponse.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	case result := <-resultChan:
		switch res := result.(type) {
		case *controllerModel.ApiError:
			controllerResponse.ResponseErrorWithApiError(c, http.StatusBadRequest, res)
			return
		case *paramCategory.AdminCreateCategoryResponse:
			controllerResponse.ResponseSuccess(c, res)
			return
		}
	}

}

// DeleteCategory 根据ID删除分类
// @Summary 根据ID删除分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param id query int true "分类ID"
// @Success 200 {object} controller.Response
// @Failure 400 {object} controller.Response
// @Failure 404 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/category [delete]
func (controller *AdminCategoryControllerImpl) DeleteCategory(c *gin.Context) {
}

// GetCategoryListByType 根据类型查询分类
// @Summary 根据类型查询分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param type query int false "分类类型：1为菜品分类，2为套餐分类"
// @Success 200 {object} controller.Response
// @Failure 400 {object} controller.Response
// @Failure 404 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/category/list [get]
func (controller *AdminCategoryControllerImpl) GetCategoryListByType(c *gin.Context) {
}
