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
	helper.ParseAndValidate[*model.RegisterRequest](ctx, &request, controller.Validate)
	data := controller.Service.Register(request)

	return ctx.Status(201).JSON(model.SuccessResponse[model.UserResponse]{
		Code:   201,
		Status: "CREATED",
		Data:   data,
	})
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	request := model.LoginRequest{}
	helper.ParseAndValidate[*model.LoginRequest](ctx, &request, controller.Validate)
	data := controller.Service.Login(request)

	return ctx.Status(200).JSON(model.SuccessResponse[model.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}
