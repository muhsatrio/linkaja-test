package account

import (
	"linkaja-test/interactors"
	"linkaja-test/platform"
)

type RequestTransfer struct {
	SenderAccountNumber   uint
	ReceiverAccountNumber uint
	Amount                int
}

func (i Interactors) Transfer(req RequestTransfer) (interactorErr interactors.Error) {
	err := i.AccountRepo.UpdateBalance(req.SenderAccountNumber, req.ReceiverAccountNumber, req.Amount)
	if err != nil {
		if err == platform.ErrInvalidInput {
			interactorErr = interactors.ErrAmoutShouldNotBeNegative
		} else if err == platform.ErrAccountNotFound {
			interactorErr = interactors.ErrAccountNotFound
		} else if err == platform.ErrInsufficientBalance {
			interactorErr = interactors.ErrInsufficientBalance
		} else if err == platform.ErrNotAllowedSameUser {
			interactorErr = interactors.ErrSendToUserItself
		} else {
			interactorErr = interactors.InternalErrorCustom(err.Error())
		}
	}

	return
}
