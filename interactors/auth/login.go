package auth

import "github.com/muhsatrio/golang-boilerplate/interactors"

type RequestLogin struct {
	Username string
	Password string
}

type ResponseLogin struct {
	Token string
}

func (s Service) Login(req RequestLogin) (resp ResponseLogin, interactorErr interactors.Error) {
	panic("implement me")
}
