package exception

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ValidationError struct {
	Code    int
	Message string
}

func (err *ValidationError) Error() string {
	return err.Message
}

func NewValidationError(err error) *fiber.Error {
	var errorStack []string

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			var param string

			if fieldError.Param() != "" {
				param = fmt.Sprintf("=%s", fieldError.Param())
			}

			message := fmt.Sprintf("error on field %s: %s%s", fieldError.Field(), fieldError.Tag(), param)
			errorStack = append(errorStack, message)
		}
	}

	return &fiber.Error{
		Code:    400,
		Message: strings.Join(errorStack, ", "),
	}
}
