package user

type CreateUserRequest struct {
	Code      string `json:"code" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Username  string `json:"userName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	RoleID    string `json:"roleId" validate:"required,uuid"`
	Status    string `json:"status"`
}
