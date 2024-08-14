package controller

import "github.com/gin-gonic/gin"

type SetMealControllerInterface interface {
	// GetSetMealList 获取套餐列表
	GetSetMealList(c *gin.Context)

	// GetSetMealDetail 获取套餐详情
	GetSetMealDetail(c *gin.Context)
}
