package mysql

//go:generate mockgen -destination=../../../mocks/user/persistence_mock.go -package=user_mock -source=user.go

import (
	"financial-planner-be/domain"

	"gorm.io/gorm"
)

// UserPersistence contains the list of functions for database table users
type UserAdapter interface {
	Create(username, password string) (user domain.User, err error)
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

func (u userRepo) Create(username, password string) (user domain.User, err error) {

	temp := User{
		Username: username,
		Password: password,
	}

	if err = u.db.Table("users").Select("username, password").Create(&temp).Error; err != nil {
		return
	}

	if err = u.db.Table("users").Select("id, username").First(&temp).Error; err != nil {
		return
	}

	user = domain.User{
		ID:       temp.ID,
		Username: temp.Username,
	}

	return
}
