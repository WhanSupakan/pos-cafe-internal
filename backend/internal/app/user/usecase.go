package user

import (
	"cafe-pos/internal/domain/errors"
	"cafe-pos/internal/domain/repository"
	"context"

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

	// Check if user already exists (example business rule)
	existingUser, _ := u.repo.GetUser(ctx, &repository.UserFilter{
		Email: &user.Email,
	})
	if existingUser != nil {
		return errors.NewConflict("User", "email", user.Email).
			WithMessage("อีเมลนี้ถูกใช้งานแล้ว")
	}

	existingUser, _ = u.repo.GetUser(ctx, &repository.UserFilter{
		Username: &user.Username,
	})
	if existingUser != nil {
		return errors.NewConflict("User", "username", user.Username).
			WithMessage("ชื่อผู้ใช้นี้ถูกใช้งานแล้ว")
	}

	return u.repo.CreateUser(ctx, user)
}

func (u *usecase) GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	if id == uuid.Nil {
		return nil, errors.NewBadRequest("invalid user ID").
			WithMessage("รหัสผู้ใช้ไม่ถูกต้อง")
	}

	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		// Repository already converts to domain error, add user-friendly message
		if appErr, ok := err.(*errors.AppError); ok && appErr.Code == "NOT_FOUND" {
			return nil, appErr.WithMessage("ไม่พบข้อมูลผู้ใช้")
		}
		return nil, err
	}

	return user, nil
}
