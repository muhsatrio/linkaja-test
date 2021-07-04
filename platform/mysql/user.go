package mysql

//go:generate mockgen -destination=../../../mocks/user/persistence_mock.go -package=user_mock -source=user.go

import (
	"gorm.io/gorm"
)

// UserPersistence contains the list of functions for database table users
type UserAdapter interface {
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
