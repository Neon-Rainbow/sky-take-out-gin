package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-gin/internal/utils/convert"
	"sky-take-out-gin/pkg/common/code"
	"sky-take-out-gin/pkg/common/request_handle"
	"sky-take-out-gin/pkg/common/response"
	"sky-take-out-gin/pkg/user/address_book/DTO"
	"sky-take-out-gin/pkg/user/address_book/service"
)

type AddressBookControllerImpl struct {
	service service.AddressBookServiceInterface
}

func (controller AddressBookControllerImpl) AddAddressBook(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}
	var req DTO.AddressBookRequestDTO
	err = c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusBadRequest, code.ParamError)
		return
	}

	apiError := controller.service.AddAddress(ctx, &req)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
}

func (controller AddressBookControllerImpl) GetAddressBookList(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	addressList, apiError := controller.service.GetUserAddressList(ctx)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, addressList)
	return
}

func (controller AddressBookControllerImpl) GetDefaultAddress(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	defaultAddress, apiError := controller.service.GetDefaultAddress(ctx)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}

	response.ResponseSuccess(c, defaultAddress)
	return
}

func (controller AddressBookControllerImpl) UpdateAddressBookByID(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	var req DTO.AddressBookRequestDTO
	apiError := controller.service.UpdateAddress(ctx, &req)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

func (controller AddressBookControllerImpl) DeleteAddressBookByID(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	addressID, err := convert.StringToUint(c.Query("address_id"))
	if err != nil {
		response.ResponseErrorWithMsg(c, http.StatusBadRequest, code.ParamError, fmt.Sprintf("address_id 参数错误, %s", err.Error()))
		return
	}

	apiError := controller.service.DeleteAddress(ctx, addressID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

func (controller AddressBookControllerImpl) GetAddressBookByID(c *gin.Context) {
	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	addressID, err := convert.StringToUint(c.Param("address_id"))
	if err != nil {
		response.ResponseErrorWithMsg(c, http.StatusBadRequest, code.ParamError, fmt.Sprintf("address_id 参数错误, %s", err.Error()))
		return
	}

	address, apiError := controller.service.GetAddressByID(ctx, addressID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, address)
	return
}

func (controller AddressBookControllerImpl) SetDefaultAddress(c *gin.Context) {
	type SetDefaultAddressRequest struct {
		AddressID uint `json:"address_id"`
	}

	ctx, err := request_handle.SetUserIDAndUsernameToContext(c)
	if err != nil {
		response.ResponseErrorWithCode(c, http.StatusInternalServerError, code.ServerError)
		return
	}

	var req SetDefaultAddressRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.ResponseErrorWithMsg(c, http.StatusBadRequest, code.ParamError, fmt.Sprintf("参数错误, %s", err.Error()))
		return
	}
	apiError := controller.service.SetDefaultAddress(ctx, req.AddressID)
	if apiError != nil {
		response.ResponseErrorWithApiError(c, http.StatusBadRequest, apiError)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

func NewAddressBookControllerImpl(service service.AddressBookServiceInterface) AddressBookControllerImpl {
	return AddressBookControllerImpl{service: service}
}
