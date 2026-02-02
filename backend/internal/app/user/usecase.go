package user

import (
	"cafe-pos/internal/domain/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *repository.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error)
}

type usecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &usecase{repo: repo}
}

func (u *usecase) CreateUser(ctx context.Context, user *repository.User) error {
	// Business logic can be added here
	// e.g., hash password, generate code, etc.

	return u.repo.CreateUser(ctx, user)
}

func (u *usecase) GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	return u.repo.GetUserByID(ctx, id)
}
