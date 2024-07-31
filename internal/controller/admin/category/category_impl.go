package category

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/model/param/admin/category"
	"time"
)

type AdminCategoryControllerImpl struct {
}

func NewAdminCategoryControllerImpl() *AdminCategoryControllerImpl {
	return &AdminCategoryControllerImpl{}
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
// @Success 200 {object} controller.Response
// @Failure 400 {object} controller.Response
// @Failure 404 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/category [put]
func (controller *AdminCategoryControllerImpl) UpdateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	go func() {
		var req category.AdminUpdateCategoryRequest
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {

		}
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {

		}
		if errors.Is(ctx.Err(), context.Canceled) {
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
