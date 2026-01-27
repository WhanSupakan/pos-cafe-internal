package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"userName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RoleID    uuid.UUID `json:"roleId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserFilter struct {
	ID       *uuid.UUID
	Code     *string
	Username *string
	Email    *string
	Status   *string
	RoleID   *uuid.UUID
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error

	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUser(ctx context.Context, filter *UserFilter) (*User, error)
	GetUsers(ctx context.Context, filter *UserFilter) ([]*User, error)

	UpdateUser(ctx context.Context, u *User) error
	UpdateUserWithFields(ctx context.Context, u *User, fields []string, omitFields []string) error
	UpdateUsers(ctx context.Context, users []*User) error
	UpdateUsersWithFields(ctx context.Context, users []*User, fields []string, omitFields []string) error

	DeleteUser(ctx context.Context, id uuid.UUID) error
}
