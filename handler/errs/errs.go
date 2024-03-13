package errs

import (
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code      int
	Messesage string
}

func (e AppError) Error() string {
	return e.Messesage
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:      fiber.StatusNotFound,
		Messesage: message,
	}
}

func NewUnexpectedError(message string) error {
	return AppError{
		Code:      fiber.StatusInternalServerError,
		Messesage: message,
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:      fiber.StatusUnprocessableEntity,
		Messesage: "Validation error",
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:      fiber.StatusBadRequest,
		Messesage: message,
	}
}

func NewUnauthorizedError(message string) error {
	return AppError{
		Code:      fiber.StatusUnauthorized,
		Messesage: message,
	}
}

func NewForbiddenError(message string) error {
	return AppError{
		Code:      fiber.StatusForbidden,
		Messesage: message,
	}
}

func NewBadgatewayError(message string) error {
	return AppError{
		Code:      fiber.StatusBadGateway,
		Messesage: message,
	}
}

func NewConflictError(message string) error {
	return AppError{
		Code:      fiber.StatusConflict,
		Messesage: message,
	}
}

func NewInternalServerError(message string) error {
	return AppError{
		Code:      fiber.StatusInternalServerError,
		Messesage: message,
	}
}
