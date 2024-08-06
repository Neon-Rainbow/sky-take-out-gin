package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `gorm:"type:varchar(32);not null;comment:'姓名'"`
	Username   string `gorm:"type:varchar(32);not null;unique;comment:'用户名'"`
	Password   string `gorm:"type:varchar(64);not null;comment:'密码'"`
	Phone      string `gorm:"type:varchar(11);not null;comment:'手机号'"`
	Sex        string `gorm:"type:varchar(2);comment:'性别'"`
	IDNumber   string `gorm:"type:varchar(18);not null;comment:'身份证号'"`
	Status     int    `gorm:"default:1;comment:'状态 0:禁用，1:启用'"`
	CreateUser uint   `gorm:"comment:'创建人'"`
	UpdateUser uint   `gorm:"comment:'修改人'"`
}
