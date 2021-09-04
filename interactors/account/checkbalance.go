package account

import (
	"linkaja-test/interactors"
	"linkaja-test/platform"
)

type ResponseCheckBalance struct {
	AccountNumber uint
	CustomerName  string
	Balance       int
}

func (i Interactors) CheckBalance(accountNumber uint) (resp ResponseCheckBalance, interactorErr interactors.Error) {
	user, err := i.accountRepo.Get(accountNumber)
	if err != nil {
		if err == platform.ErrAccountNotFound {
			interactorErr = interactors.ErrAccountNotFound
		} else {
			interactorErr = interactors.InternalErrorCustom(err.Error())
		}
		return
	}

	resp = ResponseCheckBalance{
		AccountNumber: user.AccountNumber,
		CustomerName:  user.CustomerName,
		Balance:       user.Balance,
	}
	return
}
