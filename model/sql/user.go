package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `copier:"-"`
	//OpenID   string `gorm:"type:varchar(45);comment:'微信用户唯一标识'"`
	Name     string `gorm:"type:varchar(32);comment:'姓名'"`
	Username string `gorm:"type:varchar(32);comment:'用户名'"`
	Password string `gorm:"type:varchar(500);comment:'密码'"`
	Phone    string `gorm:"type:varchar(11);comment:'手机号'"`
	Sex      string `gorm:"type:varchar(2);comment:'性别'"`
	IDNumber string `gorm:"type:varchar(18);comment:'身份证号'"`
	Avatar   string `gorm:"type:varchar(500);comment:'头像'"`
}
