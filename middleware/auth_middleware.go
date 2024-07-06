package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rafikhairan/academia/auth"
	"github.com/rafikhairan/academia/config"
	"strings"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization", "")
	if !strings.Contains(authHeader, "Bearer") {
		return fiber.NewError(401, "your action is unauthorized")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, err := jwt.ParseWithClaims(tokenString, &auth.UserClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(401, "signing method invalid")
		}

		return []byte(config.AppConfig.JWT.SignatureKey), nil
	})
	if err != nil {
		return fiber.NewError(401, "invalid token")
	}

	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok || !token.Valid {
		return fiber.NewError(401, "invalid token")
	}

	ctx.Locals("user", claims)

	return ctx.Next()
}
