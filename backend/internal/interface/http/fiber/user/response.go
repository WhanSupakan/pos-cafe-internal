package user

import (
	"cafe-pos/internal/domain/repository"
	"time"

	"github.com/google/uuid"
)

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"userName"`
	Email     string    `json:"email"`
	RoleID    uuid.UUID `json:"roleId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToCreateUserResponse(user *repository.User) *CreateUserResponse {
	if user == nil {
		return nil
	}
	return &CreateUserResponse{
		ID:        user.ID,
		Code:      user.Code,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		RoleID:    user.RoleID,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
