package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/authentication"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"github.com/rafikhairan/academia/service"
)

type UserController struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserController(service *service.UserService, validate *validator.Validate) *UserController {
	return &UserController{
		service:  service,
		validate: validate,
	}
}

func (controller *UserController) Register(ctx *fiber.Ctx) error {
	request := model.RegisterRequest{}

	if err := helper.ParseAndValidate(ctx, &request, controller.validate); err != nil {
		return err
	}

	user, err := controller.service.Register(request)
	if err != nil {
		return err
	}

	response := model.ToWebResponse[model.UserResponse](200, user.ToUserResponse)

	return ctx.Status(201).JSON(response)
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	request := model.LoginRequest{}

	if err := helper.ParseAndValidate(ctx, &request, controller.validate); err != nil {
		return err
	}

	user, err := controller.service.Login(request)
	if err != nil {
		return err
	}

	response := model.ToWebResponse[model.UserResponse](200, user.ToUserResponse)
	response.Token = authentication.GenerateToken(user)

	return ctx.Status(200).JSON(response)
}
