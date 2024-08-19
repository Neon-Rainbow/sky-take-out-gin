package model

import (
	"fmt"
	"gorm.io/gorm"
)

// AddressBook 地址簿
type AddressBook struct {
	gorm.Model   `copier:"-"` // gorm.Model 会自动添加ID、CreatedAt、UpdatedAt、DeletedAt字段
	UserID       uint         `gorm:"not null;comment:'用户id'"`
	Consignee    string       `gorm:"type:varchar(50);comment:'收货人'"`
	Sex          string       `gorm:"type:varchar(2);comment:'性别'"`
	Phone        string       `gorm:"type:varchar(11);not null;comment:'手机号'"`
	ProvinceCode string       `gorm:"type:varchar(12);comment:'省级区划编号'"`
	ProvinceName string       `gorm:"type:varchar(32);comment:'省级名称'"`
	CityCode     string       `gorm:"type:varchar(12);comment:'市级区划编号'"`
	CityName     string       `gorm:"type:varchar(32);comment:'市级名称'"`
	DistrictCode string       `gorm:"type:varchar(12);comment:'区级区划编号'"`
	DistrictName string       `gorm:"type:varchar(32);comment:'区级名称'"`
	Detail       string       `gorm:"type:varchar(200);comment:'详细地址'"`
	Label        string       `gorm:"type:varchar(100);comment:'标签'"`
	IsDefault    bool         `gorm:"default:false;comment:'默认 0 否 1是'"`
}

// GetDetailAddress 获取详细地址
func (a *AddressBook) GetDetailAddress() string {
	return fmt.Sprintf("%s%s%s%s", a.ProvinceName, a.CityName, a.DistrictName, a.Detail)
}
