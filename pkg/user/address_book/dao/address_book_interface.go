package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type AddressBookDaoInterface interface {
	AddAddress(ctx context.Context, book *model.AddressBook) error
	GetUserAddressList(ctx context.Context, userID uint) ([]model.AddressBook, error)
	GetDefaultAddress(ctx context.Context, userID uint) (*model.AddressBook, error)
	UpdateAddress(ctx context.Context, book *model.AddressBook) error
	DeleteAddress(ctx context.Context, addressID uint) error
	GetAddressByID(ctx context.Context, addressID uint) (*model.AddressBook, error)
	SetDefaultAddress(ctx context.Context, userID uint, addressID uint) error
}
