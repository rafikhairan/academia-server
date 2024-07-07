package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/controller"
	"github.com/rafikhairan/academia/middleware"
	"github.com/rafikhairan/academia/service"
	"gorm.io/gorm"
)

func NewSubjectRoutes(db *gorm.DB, validate *validator.Validate) *fiber.App {
	route := fiber.New()

	subjectService := service.NewSubjectService(db)
	subjectController := controller.NewSubjectController(subjectService, validate)

	route.Use(middleware.AuthMiddleware)

	route.Post("/", subjectController.Create)
	route.Get("/", subjectController.GetAll)
	route.Get("/:id", subjectController.GetOne)
	route.Put("/:id", subjectController.Update)
	route.Delete("/:id", subjectController.Delete)

	return route
}
