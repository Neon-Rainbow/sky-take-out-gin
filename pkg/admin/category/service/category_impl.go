package service

import (
	"sky-take-out-gin/pkg/admin/category/dao"
	"sky-take-out-gin/pkg/common/cache"
)

type CategoryServiceImpl struct {
	dao.CategoryDaoInfertace
	cache cache.RedisCacheInterface
}

func NewCategoryService(categoryDao dao.CategoryDaoInfertace, cache cache.RedisCacheInterface) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryDao, cache}
}
