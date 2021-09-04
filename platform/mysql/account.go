package mysql

import (
	"linkaja-test/domain"
	"linkaja-test/platform"

	"gorm.io/gorm"
)

//go:generate mockgen -destination=../../mocks/platform/account/mock.go -package=mock_account linkaja-test/platform/mysql AccountAdapter

// UserPersistence contains the list of functions for database table users
type AccountAdapter interface {
	Get(accountNumber uint) (account domain.Account, err error)
	UpdateBalance(receiverAccountNumber uint, balance int) (err error)
}

type accountRepo struct {
	db *gorm.DB
}

// UserInit is to init the user persistence that contains data accounts
func UserInit(db *gorm.DB) AccountAdapter {
	return accountRepo{
		db: db,
	}
}

// User platform function

func (a accountRepo) Get(accountNumber uint) (account domain.Account, err error) {
	var temp AccountCustom

	if err = a.db.Table("accounts a").Joins("LEFT JOIN customers c").Where("a.account_number = ?", accountNumber).First(&temp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = platform.ErrAccountNotFound
		}
		return
	}

	account = domain.Account{
		AccountNumber: temp.AccountNumber,
		CustomerName:  temp.CustomerName,
		Balance:       temp.Balance,
	}
	return
}

func (a accountRepo) UpdateBalance(receiverAccountNumber uint, balance int) (err error) {
	if balance < 0 {
		err = platform.ErrInvalidInput
		return
	}

	if !(receiverAccountNumber == 555001 || receiverAccountNumber == 555002) {
		err = platform.ErrAccountNotFound
		return
	}

	senderMap := make(map[uint]uint)

	senderMap[555001] = 555002
	senderMap[555002] = 555001

	senderAccountNumber := senderMap[receiverAccountNumber]

	var sender Account

	if err = a.db.Table("accounts").Where("account_number = ?", senderAccountNumber).First(&sender).Error; err != nil {
		return
	}

	if sender.Balance-balance < 0 {
		err = platform.ErrInsufficientBalance
		return
	}

	var receiver Account

	if err = a.db.Table("accounts").Where("account_number = ?", senderAccountNumber).First(&receiver).Error; err != nil {
		return
	}

	sender.Balance = sender.Balance - balance

	receiver.Balance = receiver.Balance + balance

	if err = a.db.Table("accounts").Where("account_number = ?", senderAccountNumber).Updates(sender).Error; err != nil {
		return
	}

	if err = a.db.Table("accounts").Where("account_number = ?", receiverAccountNumber).Updates(sender).Error; err != nil {
		return
	}

	return
}
