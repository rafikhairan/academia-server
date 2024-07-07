package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rafikhairan/academia/config"
	"github.com/rafikhairan/academia/exception"
	"strings"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization", "")
	if !strings.Contains(authHeader, "Bearer") {
		panic(exception.NewUnauthorizedError("Token is required"))
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(config.AppConfig.JWT.SignatureKey), nil
	})
	if err != nil || !token.Valid {
		panic(exception.NewUnauthorizedError("Token is invalid"))
	}

	claims := token.Claims.(jwt.MapClaims)
	
	ctx.Locals("user", claims)

	return ctx.Next()
}
