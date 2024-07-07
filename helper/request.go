package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidate[T any](ctx *fiber.Ctx, request T, validate *validator.Validate) {
	err := ctx.BodyParser(request)
	PanicIfError(err)

	err = validate.Struct(request)
	PanicIfError(err)
}
