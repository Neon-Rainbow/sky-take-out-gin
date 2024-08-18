package service

import (
	"context"
	model "sky-take-out-gin/model/sql"
	apiErrorModel "sky-take-out-gin/pkg/common/api_error"
	"sky-take-out-gin/pkg/user/address_book/DTO"
)

type AddressBookServiceInterface interface {
	AddAddress(ctx context.Context, req *DTO.AddressBookRequestDTO) *apiErrorModel.ApiError
	GetUserAddressList(ctx context.Context) ([]model.AddressBook, *apiErrorModel.ApiError)
	GetDefaultAddress(ctx context.Context) (*model.AddressBook, *apiErrorModel.ApiError)
	UpdateAddress(ctx context.Context, req *DTO.AddressBookRequestDTO) *apiErrorModel.ApiError
	DeleteAddress(ctx context.Context, addressID uint) *apiErrorModel.ApiError
	GetAddressByID(ctx context.Context, addressID uint) (*model.AddressBook, *apiErrorModel.ApiError)
	SetDefaultAddress(ctx context.Context, addressID uint) *apiErrorModel.ApiError
}
