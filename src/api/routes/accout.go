package router

import (
	"github.com/gofiber/fiber/v2"
	enterHandler "github.com/notAlyosha/quiz-go/api/handlers/accout"
)

func SetupAccountRouter(api fiber.Router) {
	api.Post("/enter/signIn", enterHandler.SignIn)
	api.Post("/enter/signUp", enterHandler.SignUp)
}
