package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/controller"
	"github.com/rafikhairan/academia/service"
	"gorm.io/gorm"
)

func NewUserRoutes(db *gorm.DB, validate *validator.Validate) *fiber.App {
	route := fiber.New()

	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService, validate)

	route.Post("/register", userController.Register)
	route.Post("/login", userController.Login)

	return route
}
