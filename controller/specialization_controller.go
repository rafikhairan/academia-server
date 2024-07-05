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
	request := model.SpecializationRequest{}

	if err := helper.ParseAndValidate[*model.SpecializationRequest](ctx, &request, controller.Validate); err != nil {
		return err
	}

	specialization, err := controller.Service.Create(request)
	if err != nil {
		return err
	}

	response := model.WebResponse[model.SpecializationData]{
		Code: 201,
		Data: specialization,
	}

	return ctx.Status(response.Code).JSON(response)
}

func (controller *SpecializationController) GetAll(ctx *fiber.Ctx) error {
	specializations, err := controller.Service.GetAll()
	if err != nil {
		return err
	}

	response := model.WebResponse[[]model.SpecializationData]{
		Code: 200,
		Data: specializations,
	}

	return ctx.Status(response.Code).JSON(response)
}

func (controller *SpecializationController) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	specialization, err := controller.Service.GetOne(id)
	if err != nil {
		return err
	}

	response := model.WebResponse[model.SpecializationData]{
		Code: 200,
		Data: specialization,
	}

	return ctx.Status(response.Code).JSON(response)
}

func (controller *SpecializationController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := model.SpecializationRequest{}

	if err := helper.ParseAndValidate[*model.SpecializationRequest](ctx, &request, controller.Validate); err != nil {
		return err
	}

	specialization, err := controller.Service.Update(id, request)
	if err != nil {
		return err
	}

	response := model.WebResponse[model.SpecializationData]{
		Code: 200,
		Data: specialization,
	}

	return ctx.Status(response.Code).JSON(response)
}

func (controller *SpecializationController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	specialization, err := controller.Service.Delete(id)
	if err != nil {
		return err
	}

	response := model.WebResponse[model.SpecializationData]{
		Code: 200,
		Data: specialization,
	}

	return ctx.Status(response.Code).JSON(response)
}
