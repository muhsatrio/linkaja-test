package user

import (
	"financial-planner-be/domain"
	"financial-planner-be/interactors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
	SaltHash int
}

func (s Service) Register(req RegisterRequest) (user domain.User, serviceErr interactors.Error) {
	if req.Email == "" || req.Password == "" {
		serviceErr = interactors.ErrRequiredFieldEmpty
		return
	}

	if !isEmailValid(req.Email) {
		serviceErr = interactors.ErrInvalidInput
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

	hashPassInByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), req.SaltHash)
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

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
