package models

import "github.com/google/uuid"

// UserFilter represents filter criteria for querying users
type UserFilter struct {
	ID       *uuid.UUID
	Code     *string
	Username *string
	Email    *string
	Status   *string
	RoleID   *uuid.UUID
}
