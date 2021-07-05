package mysql

//go:generate mockgen -destination=../../../mocks/user/persistence_mock.go -package=user_mock -source=user.go

import (
	"financial-planner-be/domain"

	"gorm.io/gorm"
)

// UserPersistence contains the list of functions for database table users
type UserAdapter interface {
	Create(email, password, name string) (user domain.User, err error)
	IsExist(email string) (isTrue bool, err error)
}

type userRepo struct {
	db *gorm.DB
}

// UserInit is to init the user persistence that contains data accounts
func UserInit(db *gorm.DB) UserAdapter {
	return userRepo{
		db: db,
	}
}

// User platform function

func (u userRepo) Create(email, password, name string) (user domain.User, err error) {

	temp := User{
		Email:    email,
		Password: password,
		Name:     name,
	}

	if err = u.db.Table("users").
		Create(&temp).
		Select("id, email, name").Where("email = ?", email).First(&temp).
		Error; err != nil {
		return
	}

	user = domain.User{
		ID:    temp.ID,
		Email: temp.Email,
		Name:  temp.Name,
	}

	return
}

func (u userRepo) IsExist(email string) (isTrue bool, err error) {
	var count int64

	if err = u.db.Table("users").
		Where("email = ?", email).Count(&count).
		Error; err != nil {
		return
	}

	isTrue = int(count) > 0

	return
}
