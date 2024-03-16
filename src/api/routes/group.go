package router

import (
	"github.com/gofiber/fiber/v2"
	groupHandler "github.com/notAlyosha/quiz-go/api/handlers/group"
)

func SetupGroupRouter(api fiber.Router) {
	api.Post("/group/add", groupHandler.Create)
	api.Patch("/group/update/:fid", groupHandler.Update)
	api.Delete("/group/delete/:fid", groupHandler.Delete)

	api.Get("/group", groupHandler.GetAll)
	api.Get("/group/:fid", groupHandler.GetById)
}
