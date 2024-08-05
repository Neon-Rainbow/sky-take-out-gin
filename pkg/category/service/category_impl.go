package service

import (
	"sky-take-out-gin/pkg/category/dao"
)

type CategoryServiceImpl struct {
	dao.CategoryDaoInfertace
}

func NewCategoryService(categoryDao dao.CategoryDaoInfertace) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryDao}
}
