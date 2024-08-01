package category

import (
	"sky-take-out-gin/internal/dao"
)

type CategoryServiceImpl struct {
	dao.CategoryDaoImpl
}

func NewCategoryService(categoryDao dao.CategoryDaoImpl) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryDao}
}
