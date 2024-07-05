package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/controller"
	"github.com/rafikhairan/academia/service"
	"gorm.io/gorm"
)

func NewCourseRoutes(db *gorm.DB, validate *validator.Validate) *fiber.App {
	route := fiber.New()

	courseService := service.NewCourseService(db)
	courseController := controller.NewCourseController(courseService, validate)

	route.Post("/", courseController.Create)

	return route
}
