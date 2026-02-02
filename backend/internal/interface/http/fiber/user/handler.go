package user

import (
	"cafe-pos/internal/app/user"
	"cafe-pos/internal/domain/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	usecase   user.UserUsecase
	validator *validator.Validate
}

func NewHandler(usecase user.UserUsecase) *Handler {
	return &Handler{
		usecase:   usecase,
		validator: validator.New(),
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	// Validate request
	if err := h.validator.Struct(&req); err != nil {
		var validationErrors []string
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErr {
				validationErrors = append(validationErrors, getValidationErrorMessage(fieldErr))
			}
		} else {
			validationErrors = append(validationErrors, err.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"message": validationErrors,
		})
	}

	// Convert request to domain model
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid role ID",
			"message": err.Error(),
		})
	}

	domainUser := &repository.User{
		ID:        uuid.New(),
		Code:      req.Code,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		RoleID:    roleID,
		Status:    getStatus(req.Status),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Call usecase
	if err := h.usecase.CreateUser(c.Context(), domainUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
	}

	// Return response
	response := ToCreateUserResponse(domainUser)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid user ID format",
			"message": err.Error(),
		})
	}

	domainUser, err := h.usecase.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "User not found",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(domainUser)
}

func getValidationErrorMessage(fieldErr validator.FieldError) string {
	field := fieldErr.Field()
	tag := fieldErr.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " must be at least " + fieldErr.Param() + " characters"
	case "uuid":
		return field + " must be a valid UUID"
	default:
		return field + " is invalid"
	}
}

func getStatus(status string) string {
	if status == "" {
		return "ACTIVE"
	}
	return status
}
