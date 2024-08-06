package dao

import (
	"context"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/database"
)

type AddressBookDaoImpl struct {
	db database.DatabaseInterface
}

func (dao *AddressBookDaoImpl) AddAddress(ctx context.Context, book *model.AddressBook) error {
	if book.UserID == 0 {
		book.UserID = ctx.Value("userID").(uint)
	}
	return dao.db.GetDB().WithContext(ctx).Create(book).Error
}

func (dao *AddressBookDaoImpl) GetUserAddressList(ctx context.Context, userID uint) ([]model.AddressBook, error) {
	var addressList []model.AddressBook
	err := dao.db.GetDB().WithContext(ctx).Where("user_id = ?", userID).Find(&addressList).Error
	if err != nil {
		return nil, err
	}
	return addressList, nil
}

func (dao *AddressBookDaoImpl) GetDefaultAddress(ctx context.Context, userID uint) (*model.AddressBook, error) {
	var address model.AddressBook
	err := dao.db.GetDB().WithContext(ctx).Where("is_default = ? AND user_id = ?", true, userID).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (dao *AddressBookDaoImpl) UpdateAddress(ctx context.Context, book *model.AddressBook) error {
	if book.UserID == 0 {
		book.UserID = ctx.Value("userID").(uint)
	}
	return dao.db.GetDB().WithContext(ctx).Model(book).Updates(book).Error
}

func (dao *AddressBookDaoImpl) DeleteAddress(ctx context.Context, addressID uint) error {
	return dao.db.GetDB().WithContext(ctx).Delete(&model.AddressBook{}, addressID).Error
}

func (dao *AddressBookDaoImpl) GetAddressByID(ctx context.Context, addressID uint) (*model.AddressBook, error) {
	var address model.AddressBook
	err := dao.db.GetDB().WithContext(ctx).First(&address, addressID).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (dao *AddressBookDaoImpl) SetDefaultAddress(ctx context.Context, userID uint, addressID uint) error {
	tx := dao.db.GetDB().Begin() // 开启事务

	// 创建一个通道，用于接收事务结束的信号
	done := make(chan struct{})
	defer close(done)

	// 启动一个goroutine监听ctx.Done
	go func() {
		select {
		case <-ctx.Done():
			tx.Rollback() // ctx done，回滚事务
		case <-done:
			// 正常结束
		}
	}()

	// 将该用户的所有地址的 is_default 字段设为 false
	if err := tx.Model(&model.AddressBook{}).Where("user_id = ?", userID).Update("is_default", false).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// 将指定地址的 is_default 字段设为 true
	if err := tx.Model(&model.AddressBook{}).Where("id = ? AND user_id = ?", addressID, userID).Update("is_default", true).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 正常结束，发送信号给goroutine
	done <- struct{}{}

	return nil
}

func NewAddressBookDaoImpl(db database.DatabaseInterface) *AddressBookDaoImpl {
	return &AddressBookDaoImpl{db: db}
}
