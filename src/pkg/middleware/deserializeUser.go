package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/internal/config"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/pkg/utils"
)

func DeserializeUser(ctx *fiber.Ctx) error {
	var access_token string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		access_token = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("access_token") != "" {
		access_token = ctx.Cookies("access_token")
	}

	if access_token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config := config.LoadConfig()

	tokenClaims, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Get user record from database
	var user entityUser.UserResponse

	ctx.Locals("user", user)
	ctx.Locals("access_token_uuid", tokenClaims.TokenUuid)

	return ctx.Next()
}
