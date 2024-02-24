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
