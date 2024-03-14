package router

import (
	"github.com/gofiber/fiber/v2"
	groupHandler "github.com/notAlyosha/quiz-go/api/controllers/group"
)

func SetupGroupRouter(api fiber.Router) {
	api.Post("/group/add", groupHandler.Create)
	api.Post("/group/update/:id", groupHandler.Update)
	api.Post("/group/delete/:id", groupHandler.Delete)

	api.Get("/group/", groupHandler.GetAll)
	api.Get("/group/:id", groupHandler.GetById)
}
