package user

//go:generate mockgen -destination=../../../mocks/user/usecase_mock.go -package=user_mock -source=initiator.go

import (
	"context"

	"financial-planner-be/internal/constant/model"
	"financial-planner-be/internal/repository"
	"financial-planner-be/internal/storage/persistence"

	"github.com/iDevoid/cptx"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	Registration(ctx context.Context, user *model.User) error
}

type service struct {
	transaction    cptx.Transaction
	userRepo       repository.UserRepository
	userPersist    persistence.UserPersistence
	profilePersist persistence.ProfilePersistence
}

// Initialize takes all necessary service for domain user to run the business logic of domain user
func Initialize(
	transaction cptx.Transaction,
	userRepo repository.UserRepository,
	userPersist persistence.UserPersistence,
	profilePersist persistence.ProfilePersistence,
) Usecase {
	return &service{
		transaction,
		userRepo,
		userPersist,
		profilePersist,
	}
}
