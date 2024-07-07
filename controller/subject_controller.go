package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"github.com/rafikhairan/academia/service"
)

type SubjectController struct {
	Service  *service.SubjectService
	Validate *validator.Validate
}

func NewSubjectController(service *service.SubjectService, validate *validator.Validate) *SubjectController {
	return &SubjectController{
		Service:  service,
		Validate: validate,
	}
}

func (controller *SubjectController) Create(ctx *fiber.Ctx) error {
	request := model.CreateSubjectRequest{}
	helper.ParseAndValidate[*model.CreateSubjectRequest](ctx, &request, controller.Validate)
	data := controller.Service.Create(request)

	return ctx.Status(201).JSON(model.SuccessResponse[model.SubjectResponse]{
		Code:   201,
		Status: "CREATED",
		Data:   data,
	})
}

func (controller *SubjectController) GetAll(ctx *fiber.Ctx) error {
	data := controller.Service.GetAll()

	return ctx.Status(200).JSON(model.SuccessResponse[[]model.SubjectResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *SubjectController) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := controller.Service.GetOne(id)

	return ctx.Status(200).JSON(model.SuccessResponse[model.SubjectResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *SubjectController) Update(ctx *fiber.Ctx) error {
	request := model.UpdateSubjectRequest{
		ID: ctx.Params("id"),
	}
	helper.ParseAndValidate[*model.UpdateSubjectRequest](ctx, &request, controller.Validate)
	data := controller.Service.Update(request)

	return ctx.Status(200).JSON(model.SuccessResponse[model.SubjectResponse]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *SubjectController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	controller.Service.Delete(id)

	return ctx.Status(200).JSON(model.SuccessResponse[*model.SubjectResponse]{
		Code:   200,
		Status: "OK",
	})
}
