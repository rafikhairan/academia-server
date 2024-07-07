package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"github.com/rafikhairan/academia/service"
)

type SpecializationController struct {
	Service  *service.SpecializationService
	Validate *validator.Validate
}

func NewSpecializationController(service *service.SpecializationService, validate *validator.Validate) *SpecializationController {
	return &SpecializationController{
		Service:  service,
		Validate: validate,
	}
}

func (controller *SpecializationController) Create(ctx *fiber.Ctx) error {
	request := model.CreateSpecializationRequest{}
	helper.ParseAndValidate[*model.CreateSpecializationRequest](ctx, &request, controller.Validate)
	data := controller.Service.Create(request)

	return ctx.Status(201).JSON(model.SuccessResponse[model.SpecializationResponse]{
		Code:   201,
		Status: "CREATED",
		Data:   data,
	})
}

func (controller *SpecializationController) GetAll(ctx *fiber.Ctx) error {
	data := controller.Service.GetAll()

	return ctx.Status(200).JSON(model.SuccessResponse[[]model.SpecializationResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *SpecializationController) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := controller.Service.GetOne(id)

	return ctx.Status(200).JSON(model.SuccessResponse[model.SpecializationResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *SpecializationController) Update(ctx *fiber.Ctx) error {
	request := model.UpdateSpecializationRequest{
		ID: ctx.Params("id"),
	}
	helper.ParseAndValidate[*model.UpdateSpecializationRequest](ctx, &request, controller.Validate)
	data := controller.Service.Update(request)

	return ctx.Status(200).JSON(model.SuccessResponse[model.SpecializationResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *SpecializationController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	controller.Service.Delete(id)

	return ctx.Status(200).JSON(model.SuccessResponse[*model.SpecializationResponse]{
		Code:   200,
		Status: "OK",
	})
}
