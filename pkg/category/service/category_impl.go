package category

import (
	"sky-take-out-gin/internal/dao/admin/category"
)

type CategoryServiceImpl struct {
	category.CategoryDaoInfertace
}

func NewCategoryService(categoryDao category.CategoryDaoInfertace) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryDao}
}
