package errors

import "fmt"

type AppError struct {
	Code           string            `json:"code"`
	Message        string            `json:"message,omitempty"`
	BackendMessage string            `json:"-"`
	Details        map[string]string `json:"details,omitempty"`
	Err            error             `json:"-"`
}

func (e *AppError) Error() string {
	if e.BackendMessage != "" {
		return fmt.Sprintf("%s: %s", e.Code, e.BackendMessage)
	}
	if e.Message != "" {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Code, e.Err)
	}
	return e.Code
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) WithMessage(message string) *AppError {
	e.Message = message
	return e
}

func (e *AppError) WithBackendMessage(message string) *AppError {
	e.BackendMessage = message
	return e
}

func (e *AppError) WithDetails(key, value string) *AppError {
	if e.Details == nil {
		e.Details = make(map[string]string)
	}
	e.Details[key] = value
	return e
}

func (e *AppError) WithError(err error) *AppError {
	e.Err = err
	return e
}

func New(code string) *AppError {
	return &AppError{
		Code:    code,
		Details: make(map[string]string),
	}
}

func NewNotFound(resource, identifier string) *AppError {
	return New("NOT_FOUND").
		WithBackendMessage(fmt.Sprintf("%s not found: %s", resource, identifier)).
		WithDetails("resource", resource).
		WithDetails("identifier", identifier)
}

func NewValidationFailed(backendMsg string) *AppError {
	return New("BAD_REQUEST").WithBackendMessage(backendMsg)
}

func NewConflict(resource, field, value string) *AppError {
	return New("BAD_REQUEST").
		WithBackendMessage(fmt.Sprintf("%s with %s '%s' already exists", resource, field, value)).
		WithDetails("resource", resource).
		WithDetails("field", field).
		WithDetails("value", value)
}

func NewInternal(backendMsg string) *AppError {
	return New("INTERNAL_ERROR").WithBackendMessage(backendMsg)
}

func NewUnauthorized(backendMsg string) *AppError {
	return New("UNAUTHORIZED").WithBackendMessage(backendMsg)
}

func NewForbidden(backendMsg string) *AppError {
	return New("FORBIDDEN").WithBackendMessage(backendMsg)
}

func NewBadRequest(backendMsg string) *AppError {
	return New("BAD_REQUEST").WithBackendMessage(backendMsg)
}

func NewValidationErrors(validationErrors []string) *AppError {
	if len(validationErrors) == 0 {
		return New("BAD_REQUEST").WithBackendMessage("Validation failed")
	}

	backendMsg := "Validation failed: " + validationErrors[0]
	for i := 1; i < len(validationErrors); i++ {
		backendMsg += "; " + validationErrors[i]
	}

	err := New("BAD_REQUEST").WithBackendMessage(backendMsg)
	for i, msg := range validationErrors {
		err.Details[fmt.Sprintf("error_%d", i)] = msg
	}

	return err
}
