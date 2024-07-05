package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"github.com/rafikhairan/academia/service"
)

type UserController struct {
	Service  *service.UserService
	Validate *validator.Validate
}

func NewUserController(service *service.UserService, Validate *validator.Validate) *UserController {
	return &UserController{
		Service:  service,
		Validate: Validate,
	}
}

func (controller *UserController) Register(ctx *fiber.Ctx) error {
	request := model.RegisterRequest{}

	if err := helper.ParseAndValidate[*model.RegisterRequest](ctx, &request, controller.Validate); err != nil {
		return err
	}

	user, err := controller.Service.Register(request)
	if err != nil {
		return err
	}

	response := model.WebResponse[model.UserAuthenticationData]{
		Code: 201,
		Data: user,
	}

	return ctx.Status(response.Code).JSON(response)
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	request := model.LoginRequest{}

	if err := helper.ParseAndValidate[*model.LoginRequest](ctx, &request, controller.Validate); err != nil {
		return err
	}

	user, err := controller.Service.Login(request)
	if err != nil {
		return err
	}

	response := model.WebResponse[model.UserAuthenticationData]{
		Code: 200,
		Data: user,
	}

	return ctx.Status(response.Code).JSON(response)
}
