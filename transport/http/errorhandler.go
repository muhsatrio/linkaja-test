package http

import (
	"net/http"

	"linkaja-test/interactors"
)

type ErrorObj struct {
	Message string `json:"message"`
}

func errorHandler(err interactors.Error) (int, ErrorObj) {
	var errCode int

	switch err {
	case interactors.ErrAmoutShouldNotBeNegative, interactors.ErrInsufficientBalance, interactors.ErrSendToUserItself:
		{
			errCode = http.StatusBadRequest
		}
	case interactors.ErrAccountNotFound:
		{
			errCode = http.StatusNotFound
		}
	default:
		{
			errCode = http.StatusInternalServerError
		}
	}

	errObj := ErrorObj{
		Message: err.Error(),
	}

	return errCode, errObj
}
