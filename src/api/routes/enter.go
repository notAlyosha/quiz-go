package router

import (
	"github.com/gofiber/fiber/v2"
	enterHandler "github.com/notAlyosha/quiz-go/api/controllers/enter"
)

func SetupEnterRouter(api fiber.Router) {
	api.Get("/enter/signIn", enterHandler.SignIn)
	api.Post("/enter/signUp", enterHandler.SignUp)
	api.Post("/enter/registration", enterHandler.Registration)
}
