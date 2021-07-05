package user

import (
	"financial-planner-be/domain"
	"financial-planner-be/interactors"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
}

const (
	saltHash = 8
)

func (s Service) Register(req RegisterRequest) (user domain.User, serviceErr interactors.Error) {
	if req.Email == "" || req.Password == "" {
		serviceErr = interactors.ErrRequiredFieldEmpty
		return
	}

	isExist, err := s.UserRepo.IsExist(req.Email)
	if err != nil {
		serviceErr = interactors.InternalErrorCustom(err.Error())
		return
	}
	if isExist {
		serviceErr = interactors.ErrDuplicateDataAdd
		return
	}

	hashPassInByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), saltHash)
	if err != nil {
		serviceErr = interactors.InternalErrorCustom(err.Error())
		return
	}

	user, err = s.UserRepo.Create(req.Email, string(hashPassInByte), req.Name)
	if err != nil {
		serviceErr = interactors.InternalErrorCustom(err.Error())
		return
	}

	return
}
