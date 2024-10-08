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
//11:用户端分类模块错误
//具体错误码
//2 位数进行表示
//具体根据模块中的操作定义

const (
	ServerError           ResponseCode = 10101
	TooManyRequests       ResponseCode = 10102
	ParamBindError        ResponseCode = 10103
	AuthorizationError    ResponseCode = 10104
	UrlSignError          ResponseCode = 10105
	MySQLExecError        ResponseCode = 10106
	ParamError            ResponseCode = 10107
	RequestForbidden      ResponseCode = 10108
	CacheInvalidateFailed ResponseCode = 10109

	// 管理端分类模块错误
	CategoryBindParamError     ResponseCode = 20101
	CategoryNotExist           ResponseCode = 20102
	CategoryUpdateFailed       ResponseCode = 20103
	CategoryGetFailed          ResponseCode = 20104
	CategoryChangeStatusFailed ResponseCode = 20105
	CategoryCreateFailed       ResponseCode = 20106
	CategoryDeleteFailed       ResponseCode = 20107

	// 管理端员工模块错误
	EmployeeBindParamError     ResponseCode = 20201
	EmployeeEditPasswordFailed ResponseCode = 20202
	EmployeeNotFound           ResponseCode = 20203
	EmployeeChangeStatusFailed ResponseCode = 20204
	EmployeeGetPageFailed      ResponseCode = 20205
	EmployeeLoginFailed        ResponseCode = 20206
	EmployeeSearchFailed       ResponseCode = 20207
	EmployeeEditFailed         ResponseCode = 20208
	EmployeeAddFailed          ResponseCode = 20209

	// 套餐模块错误
	UpdateSetmealError  ResponseCode = 20301
	GetSetmealPageError ResponseCode = 20302
	CreateSetmealError  ResponseCode = 20303

	// 商店模块错误
	GetShopStatusError ResponseCode = 20401
	SetShopStatusError ResponseCode = 20402

	// 菜品模块错误
	UpdateDishError           ResponseCode = 20501
	DeleteDishError           ResponseCode = 20502
	CreateDishError           ResponseCode = 20503
	SearchDishByIDError       ResponseCode = 20504
	SearchDishByCategoryError ResponseCode = 20505
	SearchDishByPageError     ResponseCode = 20506
	ChangeDishStatusError     ResponseCode = 20507
	UpdateDishFlavorError     ResponseCode = 20508
	CreateDishFlavorError     ResponseCode = 20509
	DeleteDishFlavorError     ResponseCode = 20510

	// 用户端分类模块错误
	CategoryGetListError ResponseCode = 21101

	//地址模块错误
	AddAddressError         ResponseCode = 21201
	GetUserAddressListError ResponseCode = 21202
	GetDefaultAddressError  ResponseCode = 21203
	UpdateAddressError      ResponseCode = 21204
	DeleteAddressError      ResponseCode = 21205
	GetAddressByIDError     ResponseCode = 21206

	// 套餐模块错误
	SetMealGetListError   ResponseCode = 21301
	SetMealGetDetailError ResponseCode = 21302

	// 用户端登录模块
	UserLoginError        ResponseCode = 21401
	UserLogoutError       ResponseCode = 21402
	UserRefreshTokenError ResponseCode = 21403
	UserRegisterError     ResponseCode = 21404

	// 菜品浏览模块
	GetDishByIdError ResponseCode = 21501

	// 用户端订单模块
	GetOrderDetailError       ResponseCode = 21601 // 获取订单详情失败
	SubmitOrderError          ResponseCode = 21602 // 提交订单失败
	GetHistoryOrderError      ResponseCode = 21603 // 获取历史订单失败
	CancelOrderError          ResponseCode = 21604 // 取消订单失败
	PayOrderError             ResponseCode = 21605 // 支付订单失败
	FinishOrderError          ResponseCode = 21606 // 完成订单失败
	RejectOrderError          ResponseCode = 21607 // 拒绝订单失败
	ConfirmOrderError         ResponseCode = 21608 // 确认订单失败
	DeliveryOrderError        ResponseCode = 21609 // 发货订单失败
	ConditionSearchOrderError ResponseCode = 21610 // 条件查询订单失败

	// 购物车模块
	AddCartToCacheError              ResponseCode = 21701
	GetCartListFromCacheError        ResponseCode = 21702
	DeleteCartFromCacheError         ResponseCode = 21703
	GetCartTotalAmountFromRedisError ResponseCode = 21704
	SaveRedisToDatabaseError         ResponseCode = 21705

	// 权限校验失败
	RequestUnauthorized ResponseCode = 21001
)

var ResponseCodeMessageMap = map[ResponseCode]string{
	ServerError:        "服务器内部错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数信息错误",
	AuthorizationError: "签名信息错误",
	UrlSignError:       "参数签名错误",
	MySQLExecError:     "数据库执行错误",
	ParamError:         "参数错误",
	RequestForbidden:   "请求被拒绝",

	//管理端分类模块错误
	CategoryBindParamError:     "分类参数绑定错误",
	CategoryNotExist:           "分类不存在",
	CategoryUpdateFailed:       "更新分类失败",
	CategoryGetFailed:          "获取分类失败",
	CategoryChangeStatusFailed: "修改分类状态失败",
	CategoryCreateFailed:       "创建分类失败",
	CategoryDeleteFailed:       "删除分类失败",

	//管理端员工模块错误
	EmployeeBindParamError:     "员工参数绑定错误",
	EmployeeEditPasswordFailed: "修改员工密码失败",
	EmployeeNotFound:           "员工不存在",
	EmployeeChangeStatusFailed: "修改员工状态失败",
	EmployeeGetPageFailed:      "获取员工分页失败",
	EmployeeLoginFailed:        "员工登录失败",
	EmployeeSearchFailed:       "员工查询失败",
	EmployeeEditFailed:         "员工编辑失败",
	EmployeeAddFailed:          "员工添加失败",

	// 套餐模块错误
	UpdateSetmealError:  "更新套餐失败",
	GetSetmealPageError: "获取套餐分页失败",
	CreateSetmealError:  "创建套餐失败",

	// 店铺模块错误
	GetShopStatusError: "获取店铺状态失败",
	SetShopStatusError: "设置店铺状态失败",

	// 菜品模块错误
	UpdateDishError:           "更新菜品失败",
	DeleteDishError:           "删除菜品失败",
	CreateDishError:           "创建菜品失败",
	SearchDishByIDError:       "查询菜品失败",
	SearchDishByCategoryError: "查询菜品失败",
	SearchDishByPageError:     "查询菜品失败",
	ChangeDishStatusError:     "修改菜品状态失败",
	UpdateDishFlavorError:     "更新菜品口味失败",
	CreateDishFlavorError:     "创建菜品口味失败",
	DeleteDishFlavorError:     "删除菜品口味失败",

	// 权限校验失败
	RequestUnauthorized: "请求未授权",
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
