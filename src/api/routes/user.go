package router

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/notAlyosha/quiz-go/api/controllers/user"
)

func SetupUserRouter(api fiber.Router) {
	api.Post("/user/add", userHandler.Create)
	api.Post("/user/update/:id", userHandler.Update)
	api.Post("/user/delete/:id", userHandler.Delete)

	api.Get("/user", userHandler.GetAll)
	api.Get("/user/:id", userHandler.GetById)
	api.Get("/user/group/:id", userHandler.GetGroupById)
	api.Get("/user/role/:id", userHandler.GetByRoleId)

}
