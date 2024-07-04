package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/database"
	"github.com/rafikhairan/academia/middleware"
	"github.com/rafikhairan/academia/route"
)

func main() {
	db := database.NewDB()
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	validate := validator.New()
	router := route.NewRouter(db, validate)

	app.Mount("/api", router)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
