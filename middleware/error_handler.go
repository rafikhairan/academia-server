package middleware

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
	"github.com/rafikhairan/academia/exception"
	"github.com/rafikhairan/academia/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return throwValidationErrors(ctx, validationErrors)
	}

	var notFoundError *exception.NotFoundError
	if errors.As(err, &notFoundError) {
		return throwNotFoundError(ctx, notFoundError)
	}

	var badRequestError *exception.BadRequestError
	if errors.As(err, &badRequestError) {
		return throwBadRequestError(ctx, badRequestError)
	}

	var unauthorizedError *exception.UnauthorizedError
	if errors.As(err, &unauthorizedError) {
		return throwUnauthorizedError(ctx, unauthorizedError)
	}

	return throwInternalServerError(ctx, err)
}

func throwValidationErrors(ctx *fiber.Ctx, errors validator.ValidationErrors) error {
	response := model.ErrorResponse[map[string]any]{
		Code:   400,
		Status: "BAD_REQUEST",
		Errors: make(map[string]any),
	}

	for _, fieldError := range errors {
		var param string

		if fieldError.Param() != "" {
			param = fmt.Sprintf("=%s", strcase.ToSnake(fieldError.Param()))
		}

		field := strcase.ToSnake(fieldError.Field())
		message := fmt.Sprintf("must %s%s", fieldError.Tag(), param)

		response.Errors[field] = message
	}

	return ctx.Status(response.Code).JSON(response)
}

func throwInternalServerError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(500).JSON(model.ErrorResponse[string]{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Errors: err.Error(),
	})
}

func throwNotFoundError(ctx *fiber.Ctx, err *exception.NotFoundError) error {
	return ctx.Status(err.Code).JSON(model.ErrorResponse[string]{
		Code:   err.Code,
		Status: err.Status,
		Errors: err.Error(),
	})
}

func throwBadRequestError(ctx *fiber.Ctx, err *exception.BadRequestError) error {
	return ctx.Status(err.Code).JSON(model.ErrorResponse[string]{
		Code:   err.Code,
		Status: err.Status,
		Errors: err.Error(),
	})
}

func throwUnauthorizedError(ctx *fiber.Ctx, err *exception.UnauthorizedError) error {
	return ctx.Status(err.Code).JSON(model.ErrorResponse[string]{
		Code:   err.Code,
		Status: err.Status,
		Errors: err.Error(),
	})
}
