package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	response := model.WebResponse[any]{
		Code:   500,
		Errors: err.Error(),
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		response.Code = fiberError.Code
		response.Errors = fiberError.Error()
	}

	return ctx.Status(response.Code).JSON(response)
}
