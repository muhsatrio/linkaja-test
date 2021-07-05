package http

import (
	"financial-planner-be/interactors"
	"net/http"
)

type ErrorObj struct {
	Message string `json:"message"`
}

func errorHandler(err interactors.Error) (int, ErrorObj) {
	var errCode int

	switch err {
	case interactors.ErrInvalidInput, interactors.ErrRequiredFieldEmpty, interactors.ErrDuplicateDataAdd:
		{
			errCode = http.StatusBadRequest
		}
	case interactors.ErrUnauthorized:
		{
			errCode = http.StatusUnauthorized
		}
	case interactors.ErrForbiddenAccess:
		{
			errCode = http.StatusForbidden
		}
	case interactors.ErrDataNotFound:
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
