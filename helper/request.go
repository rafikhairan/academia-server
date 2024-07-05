package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/exception"
)

func ParseAndValidate[T any](ctx *fiber.Ctx, request T, validate *validator.Validate) error {
	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	if err := validate.Struct(request); err != nil {
		return exception.NewValidationError(err)
	}

	return nil
}
