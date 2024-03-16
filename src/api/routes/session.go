package router

import (
	"github.com/gofiber/fiber/v2"
	sessionHandler "github.com/notAlyosha/quiz-go/api/handlers/session"
)

func SetupSessiontRoutes(api fiber.Router) {
	api.Post("/subject/add", sessionHandler.Add)
	api.Patch("/subject/update", sessionHandler.Update)
	api.Get("/subject", sessionHandler.GetAll)
	api.Get("/subject/:fid", sessionHandler.GetByID)
	api.Get("/subject/user/:fid", sessionHandler.GetByUserID)
	api.Get("/subject/group/:fid", sessionHandler.GetByGroupID)
}
