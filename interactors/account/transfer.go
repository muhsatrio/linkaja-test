package account

import (
	"linkaja-test/interactors"
	"linkaja-test/platform"
)

type RequestTransfer struct {
	AccountNumber uint
	Amount        int
}

func (i Interactors) Transfer(req RequestTransfer) (interactorErr interactors.Error) {
	err := i.accountRepo.UpdateBalance(req.AccountNumber, req.Amount)
	if err != nil {
		if err == platform.ErrInvalidInput {
			interactorErr = interactors.ErrAmoutShouldNotBeNegative
		} else if err == platform.ErrAccountNotFound {
			interactorErr = interactors.ErrAccountNotFound
		} else if err == platform.ErrInsufficientBalance {
			interactorErr = interactors.ErrInsufficientBalance
		} else {
			interactorErr = interactors.InternalErrorCustom(err.Error())
		}
	}

	return
}
