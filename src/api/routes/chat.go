package chatRouter

import (
	"github.com/gofiber/fiber/v2"
	chatHandler "github.com/notAlyosha/quiz-go/api/controllers/chat"
)

func SetupChatRouter(api fiber.Router) {
	api.Post("/chat/add", chatHandler.Create)
	api.Post("/chat/update/:id", chatHandler.Update)
	api.Post("/chat/delete/:id", chatHandler.Delete)

	api.Get("/chat", chatHandler.GetAll)
	api.Get("/chat/:id", chatHandler.GetById)
	api.Get("/chat/group/:id", chatHandler.GetGroupById)
	api.Get("/chat/user/:id", chatHandler.GetByUserId)
	api.Get("/chat/role/:id", chatHandler.GetByRoleId)
}
