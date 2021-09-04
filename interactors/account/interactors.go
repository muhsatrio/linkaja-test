package account

import (
	"linkaja-test/interactors"
	"linkaja-test/platform/mysql"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	CheckBalance(accountNumber uint) (resp ResponseCheckBalance, interactorErr interactors.Error)
	Transfer(req RequestTransfer) (interactorErr interactors.Error)
}

type Interactors struct {
	accountRepo mysql.AccountAdapter
}

var _ Usecase = Interactors{}
