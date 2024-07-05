package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Random struct {
	Name    string `validate:"required,min=5" json:"name"`
	Address string `validate:"required" json:"address"`
}

func NewRouter(db *gorm.DB, validate *validator.Validate) *fiber.App {
	router := fiber.New()

	userRoutes := NewUserRoutes(db, validate)
	specializationRoutes := NewSpecializationRoutes(db, validate)
	courseRoutes := NewCourseRoutes(db, validate)

	router.Mount("/users", userRoutes)
	router.Mount("/specializations", specializationRoutes)
	router.Mount("/courses", courseRoutes)

	return router
}
