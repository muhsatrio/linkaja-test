package auth

import (
	"github.com/muhsatrio/golang-boilerplate/interactors"
	"github.com/muhsatrio/golang-boilerplate/platform/jwt"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	Login(req RequestLogin) (resp ResponseLogin, interactorErr interactors.Error)
}

type Service struct {
	JwtRepo jwt.JwtAdapter
}

var _ Usecase = Service{}
