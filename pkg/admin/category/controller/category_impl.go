package controller

import (
	"github.com/gin-gonic/gin"
	serviceCategory "sky-take-out-gin/pkg/admin/category/service"
	controllerResponse "sky-take-out-gin/pkg/common/request_handle"
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
	controllerResponse.HandleRequest(c, controller.CategoryService.UpdateCategory, c.ShouldBindBodyWithJSON)
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
	controllerResponse.HandleRequest(c, controller.CategoryService.GetCategoryPage, c.ShouldBindQuery)
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
	controllerResponse.HandleRequest(c, controller.CategoryService.ChangeCategoryStatus, c.ShouldBindUri, c.ShouldBindQuery)
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
	controllerResponse.HandleRequest(c, controller.CategoryService.CreateCategory, c.ShouldBindBodyWithJSON)
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
	controllerResponse.HandleRequest(c, controller.CategoryService.DeleteCategory, c.ShouldBindQuery)
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
	controllerResponse.HandleRequest(c, controller.CategoryService.GetCategoryByType, c.ShouldBindQuery)
}
