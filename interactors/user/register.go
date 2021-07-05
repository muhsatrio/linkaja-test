package user

import (
	"financial-planner-be/domain"
	"financial-planner-be/interactors"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string
	Password string
}

const (
	saltHash = 8
)

func (s Service) Register(req RegisterRequest) (user domain.User, serviceErr interactors.Error) {
	hashPassInByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), saltHash)
	if err != nil {
		serviceErr = interactors.InternalErrorCustom(err.Error())
		return
	}

	user, err = s.UserRepo.Create(req.Username, string(hashPassInByte))
	if err != nil {
		serviceErr = interactors.InternalErrorCustom(err.Error())
		return
	}

	return
}
