package router

import (
	"github.com/gofiber/fiber/v2"
	chatHandler "github.com/notAlyosha/quiz-go/api/handlers/chat"
)

func SetupChatRouter(api fiber.Router) {
	api.Post("/chat/add", chatHandler.Create)
	api.Patch("/chat/update/:fid", chatHandler.Update)
	api.Delete("/chat/delete/:fid", chatHandler.Delete)

	api.Get("/chat", chatHandler.GetAll)
	api.Get("/chat/:fid", chatHandler.GetById)
	api.Get("/chat/user/:fid", chatHandler.GetByUserId)
}
