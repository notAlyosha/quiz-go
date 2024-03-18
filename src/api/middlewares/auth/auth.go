package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/notAlyosha/quiz-go/api/config"
)

func JwtMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Token")
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing token"})
		}

		// Validate token
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return config.GetConfig().GetPrivateKey(), nil
		})
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}
		if !tkn.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token is invalid"})
		}

		return ctx.Next()
	}
}
