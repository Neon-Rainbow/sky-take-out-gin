package category

import "github.com/gin-gonic/gin"

// AdminCategoryControllerInterface 定义控制器接口
type AdminCategoryControllerInterface interface {
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryPage(c *gin.Context)
	GetCategoryListByType(c *gin.Context)
	ChangeCategoryStatus(c *gin.Context)
}
