package platform

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidInput        = errors.New("invalid input")
	ErrNotAllowedSameUser  = errors.New("not allowed to transfer to sender itself")
)
