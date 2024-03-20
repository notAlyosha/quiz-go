package accoutHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/notAlyosha/quiz-go/api/config"
)

func SignUp(ctx *fiber.Ctx) error {
	return nil
}

func Registration(ctx *fiber.Ctx) error {
	return nil
}

func SignIn(ctx *fiber.Ctx) error {
	// TODO: set claims if user exsists
	claims := jwt.MapClaims{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.GetConfig().AccessTokenPrivateKey)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error signing token"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": signedToken})
}
