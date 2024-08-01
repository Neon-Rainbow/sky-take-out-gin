// Package code 响应码定义
package code

import "strconv"

// ResponseCode 响应码
type ResponseCode int

// GenerateResponseCode 生成响应码
// @Param ServiceError int 服务级错误码：1 位数进行表示，比如 1 为系统级错误；2 为普通错误，通常是由用户非法操作引起。
// @Param ModelError int 模块级错误码：2 位数进行表示，比如 01 为用户模块错误；02 为商品模块错误。
// @Param DetailError int 具体错误码：2 位数进行表示，比如 01 为用户不存在；02 为密码错误。
// @Return ResponseCode 响应码
func GenerateResponseCode(ServiceError int, ModelError int, DetailError int) ResponseCode {
	return ResponseCode(ServiceError*10000 + ModelError*100 + DetailError)
}

//服务级错误码
//1 位数进行表示
//其中 1 为系统级错误, 2 为普通错误, 通常是由用户非法操作引起
//
//模块级错误码
//2 位数进行表示
//
//01:管理端分类模块错误
//
//具体错误码
//2 位数进行表示
//具体根据模块中的操作定义

const (
	ServerError        ResponseCode = 10101
	TooManyRequests    ResponseCode = 10102
	ParamBindError     ResponseCode = 10103
	AuthorizationError ResponseCode = 10104
	UrlSignError       ResponseCode = 10105
	MySQLExecError     ResponseCode = 10106

	// 管理端分类模块错误
	CategoryBindParamError ResponseCode = 20101
	CategoryNotExist       ResponseCode = 20102
	CategoryUpdateFailed   ResponseCode = 20103
	CategoryGetFailed      ResponseCode = 20104
)

var ResponseCodeMessageMap = map[ResponseCode]string{
	ServerError:        "服务器内部错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数信息错误",
	AuthorizationError: "签名信息错误",
	UrlSignError:       "参数签名错误",
	MySQLExecError:     "数据库执行错误",

	//管理端分类模块错误
	CategoryBindParamError: "分类参数绑定错误",
	CategoryNotExist:       "分类不存在",
	CategoryUpdateFailed:   "更新分类失败",
	CategoryGetFailed:      "获取分类失败",
}

func (r ResponseCode) Message() string {
	return ResponseCodeMessageMap[r]
}

func (r ResponseCode) String() string {
	return strconv.Itoa(int(r))
}

func (r ResponseCode) Int() int {
	return int(r)
}

func (r ResponseCode) Error() string {
	return r.Message()
}

func (r ResponseCode) GetServerErrorCode() int {
	return int(r / 10000)
}

func (r ResponseCode) GetModelErrorCode() int {
	return int(r % 10000 / 100)
}

func (r ResponseCode) GetDetailErrorCode() int {
	return int(r % 100)
}
