package auth

import (
	"github.com/muhsatrio/golang-boilerplate/domain"
	"github.com/muhsatrio/golang-boilerplate/interactors"
	"github.com/muhsatrio/golang-boilerplate/platform"
	"github.com/muhsatrio/golang-boilerplate/platform/bcrypt"
)

type RequestLogin struct {
	Email    string
	Password string
	SaltHash int
}

type ResponseLogin struct {
	Token string
}

func (s Interactors) Login(req RequestLogin) (resp ResponseLogin, interactorErr interactors.Error) {
	user, err := s.UserRepo.FindUser(req.Email)
	if err != nil {
		if err == platform.ErrNotFound {
			interactorErr = interactors.ErrUnauthorized
		} else {
			interactorErr = interactors.InternalErrorCustom(err.Error())
		}
		return
	}

	err = bcrypt.Compare(req.Password, user.Password, req.SaltHash)
	if err != nil {
		interactorErr = interactors.ErrUnauthorized
		return
	}

	claim := domain.TokenClaim{
		Email: user.Email,
		Name:  user.Name,
	}

	token, err := s.JwtRepo.GenerateToken(claim)
	if err != nil {
		interactorErr = interactors.InternalErrorCustom(err.Error())
		return
	}

	resp = ResponseLogin{
		Token: token,
	}

	return
}
