package errors

import (
	"fmt"

	"cafe-pos/internal/domain/errors"
	"cafe-pos/internal/domain/logger"
	"cafe-pos/internal/db/models"

	"github.com/gofiber/fiber/v2"
)

var statusCodeMap = map[string]int{
	"NOT_FOUND":      fiber.StatusNotFound,
	"BAD_REQUEST":    fiber.StatusBadRequest,
	"UNAUTHORIZED":   fiber.StatusUnauthorized,
	"FORBIDDEN":      fiber.StatusForbidden,
	"INTERNAL_ERROR": fiber.StatusInternalServerError,
}

type ErrorHandler struct {
	logger logger.Logger
}

func NewErrorHandler(l logger.Logger) *ErrorHandler {
	return &ErrorHandler{logger: l}
}

func (h *ErrorHandler) HandleError(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	if appErr, ok := err.(*errors.AppError); ok {
		return h.handleAppError(c, appErr)
	}

	var appErr *errors.AppError
	switch {
	case models.IsNotFound(err):
		appErr = errors.NewNotFound("Resource", "").WithError(err)
	case models.IsConstraintError(err):
		appErr = errors.NewBadRequest("Resource constraint violation").WithError(err)
	case models.IsValidationError(err):
		appErr = errors.NewBadRequest("Validation failed").WithError(err)
	default:
		appErr = errors.NewInternal("An unexpected error occurred").WithError(err)
	}

	return h.handleAppError(c, appErr)
}

func (h *ErrorHandler) handleAppError(c *fiber.Ctx, appErr *errors.AppError) error {
	if appErr.BackendMessage != "" {
		h.logger.LogError(appErr.Code, appErr.BackendMessage, appErr.Details, appErr.Err)
	}

	statusCode := getStatusCode(appErr.Code)
	response := make(fiber.Map, 3)
	response["error"] = appErr.Code

	validationErrors := extractValidationErrors(appErr.Details)
	if len(validationErrors) > 0 {
		if len(validationErrors) == 1 {
			response["message"] = validationErrors[0]
		} else {
			response["message"] = validationErrors
		}
	} else if appErr.Message != "" {
		response["message"] = appErr.Message
	}

	if len(appErr.Details) > 0 {
		details := filterDetails(appErr.Details)
		if len(details) > 0 {
			response["details"] = details
		}
	}

	return c.Status(statusCode).JSON(response)
}

func HandleError(c *fiber.Ctx, err error) error {
	return defaultHandler.HandleError(c, err)
}

var defaultHandler *ErrorHandler

func SetDefaultHandler(l logger.Logger) {
	defaultHandler = NewErrorHandler(l)
}

func extractValidationErrors(details map[string]string) []string {
	if len(details) == 0 {
		return nil
	}

	var errors []string
	for i := 0; ; i++ {
		key := fmt.Sprintf("error_%d", i)
		if msg, exists := details[key]; exists {
			errors = append(errors, msg)
		} else {
			break
		}
	}

	return errors
}

func isValidationErrorKey(key string) bool {
	const prefix = "error_"
	if len(key) <= len(prefix) || key[:len(prefix)] != prefix {
		return false
	}

	suffix := key[len(prefix):]
	if len(suffix) == 0 {
		return false
	}

	for _, r := range suffix {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func filterDetails(details map[string]string) map[string]string {
	filtered := make(map[string]string, len(details))
	for k, v := range details {
		if !isValidationErrorKey(k) {
			filtered[k] = v
		}
	}
	return filtered
}

func getStatusCode(code string) int {
	if status, ok := statusCodeMap[code]; ok {
		return status
	}
	return fiber.StatusInternalServerError
}
