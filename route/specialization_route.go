package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/controller"
	"github.com/rafikhairan/academia/service"
	"gorm.io/gorm"
)

func NewSpecializationRoutes(db *gorm.DB, validate *validator.Validate) *fiber.App {
	route := fiber.New()

	specializationService := service.NewSpecializationService(db)
	specializationController := controller.NewSpecializationController(specializationService, validate)

	route.Post("/", specializationController.Create)
	route.Get("/", specializationController.GetAll)
	route.Get("/:id", specializationController.GetOne)
	route.Put("/:id", specializationController.Update)
	route.Delete("/:id", specializationController.Delete)

	return route
}
