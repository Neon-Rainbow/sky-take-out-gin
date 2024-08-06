package controller

import "github.com/gin-gonic/gin"

type CategoryControllerInterface interface {
	// GetCategoryList 获取分类列表
	GetCategoryList(c *gin.Context)
}
