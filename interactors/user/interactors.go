package user

//go:generate mockgen -destination=../../../mocks/user/usecase_mock.go -package=user_mock -source=initiator.go

import (
	"github.com/muhsatrio/golang-boilerplate/domain"
	"github.com/muhsatrio/golang-boilerplate/interactors"
	"github.com/muhsatrio/golang-boilerplate/platform/mysql"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	Register(req RegisterRequest) (user domain.User, serviceErr interactors.Error)
}

type Interactors struct {
	UserRepo mysql.UserAdapter
}

var _ Usecase = Interactors{}
