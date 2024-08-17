package service

import (
	"context"
	"github.com/jinzhu/copier"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/code"
	apiErrorModel "sky-take-out-gin/pkg/common/error"
	"sky-take-out-gin/pkg/user/address_book/DTO"
	"sky-take-out-gin/pkg/user/address_book/dao"
)

type AddressBookServiceImpl struct {
	dao dao.AddressBookDaoInterface
}

func (service AddressBookServiceImpl) AddAddress(ctx context.Context, req *DTO.AddressBookRequestDTO) *apiErrorModel.ApiError {
	userID := ctx.Value("userID").(uint)
	req.UserID = userID
	address := &model.AddressBook{}
	err := copier.CopyWithOption(address, req, copier.Option{IgnoreEmpty: true})
	//err := convert.UpdateStructFields(req, address)
	if err != nil {
		return &apiErrorModel.ApiError{
			Code: code.AddAddressError,
			Msg:  err.Error(),
		}
	}
	err = service.dao.AddAddress(ctx, address)
	if err != nil {
		return &apiErrorModel.ApiError{
			Code: code.AddAddressError,
			Msg:  err.Error(),
		}
	}
	return nil
}

func (service AddressBookServiceImpl) GetUserAddressList(ctx context.Context) ([]model.AddressBook, *apiErrorModel.ApiError) {
	addressList, err := service.dao.GetUserAddressList(ctx, ctx.Value("userID").(uint))
	if err != nil {
		return nil, &apiErrorModel.ApiError{
			Code: code.GetUserAddressListError,
			Msg:  err.Error(),
		}
	}
	return addressList, nil
}

func (service AddressBookServiceImpl) GetDefaultAddress(ctx context.Context) (*model.AddressBook, *apiErrorModel.ApiError) {
	address, err := service.dao.GetDefaultAddress(ctx, ctx.Value("userID").(uint))
	if err != nil {
		return nil, &apiErrorModel.ApiError{
			Code: code.GetDefaultAddressError,
			Msg:  err.Error(),
		}
	}
	return address, nil
}

func (service AddressBookServiceImpl) UpdateAddress(ctx context.Context, req *DTO.AddressBookRequestDTO) *apiErrorModel.ApiError {
	book := &model.AddressBook{}
	err := copier.CopyWithOption(book, req, copier.Option{IgnoreEmpty: true})
	//err := convert.UpdateStructFields(book, req)
	if err != nil {
		return &apiErrorModel.ApiError{
			Code: code.UpdateAddressError,
			Msg:  err.Error(),
		}
	}
	err = service.dao.UpdateAddress(ctx, book)
	if err != nil {
		return &apiErrorModel.ApiError{
			Code: code.UpdateAddressError,
			Msg:  err.Error(),
		}
	}
	return nil
}

func (service AddressBookServiceImpl) DeleteAddress(ctx context.Context, addressID uint) *apiErrorModel.ApiError {
	err := service.dao.DeleteAddress(ctx, addressID)
	if err != nil {
		return &apiErrorModel.ApiError{
			Code: code.DeleteAddressError,
			Msg:  err.Error(),
		}
	}
	return nil
}

func (service AddressBookServiceImpl) GetAddressByID(ctx context.Context, addressID uint) (*model.AddressBook, *apiErrorModel.ApiError) {
	address, err := service.dao.GetAddressByID(ctx, addressID)
	if err != nil {
		return nil, &apiErrorModel.ApiError{
			Code: code.GetAddressByIDError,
			Msg:  err.Error(),
		}
	}
	return address, nil
}

func (service AddressBookServiceImpl) SetDefaultAddress(ctx context.Context, addressID uint) *apiErrorModel.ApiError {
	userID := ctx.Value("userID").(uint)
	err := service.dao.SetDefaultAddress(ctx, userID, addressID)
	if err != nil {
		return &apiErrorModel.ApiError{
			Code: code.UpdateAddressError,
			Msg:  err.Error(),
		}
	}
	return nil
}

func NewAddressBookServiceImpl(dao dao.AddressBookDaoInterface) *AddressBookServiceImpl {
	return &AddressBookServiceImpl{dao: dao}
}
