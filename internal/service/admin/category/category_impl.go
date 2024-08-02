package category

import (
	"sky-take-out-gin/internal/dao/admin/category"
)

type CategoryServiceImpl struct {
	category.CategoryDao
}

func NewCategoryService(categoryDao category.CategoryDao) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryDao}
}
