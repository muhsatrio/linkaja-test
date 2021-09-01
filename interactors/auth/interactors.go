package auth

import (
	"github.com/muhsatrio/golang-boilerplate/interactors"
	"github.com/muhsatrio/golang-boilerplate/platform/jwt"
	"github.com/muhsatrio/golang-boilerplate/platform/mysql"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	Login(req RequestLogin) (resp ResponseLogin, interactorErr interactors.Error)
}

type Interactors struct {
	JwtRepo  jwt.JwtAdapter
	UserRepo mysql.UserAdapter
}

var _ Usecase = Interactors{}
