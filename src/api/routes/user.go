package router

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/notAlyosha/quiz-go/api/controllers/user"
)

func SetupUserRouter(api fiber.Router) {
	api.Post("/user/add", userHandler.Create)
	api.Patch("/user/update/:fid", userHandler.Update)
	api.Delete("/user/delete/:fid", userHandler.Delete)

	api.Get("/user", userHandler.GetAll)
	api.Get("/user/:fid", userHandler.GetById)
	api.Get("/user/group/:fid", userHandler.GetGroupById)
	api.Get("/user/role/:fid", userHandler.GetByRoleId)

}
