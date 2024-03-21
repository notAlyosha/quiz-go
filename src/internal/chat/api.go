package chat

import (
	"github.com/gofiber/fiber/v2"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"

	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupChatRouter(api fiber.Router) {
	api.Post("/chat/add", middleware.DeserializeUser, create)
	api.Patch("/chat/update/:fid", middleware.DeserializeUser, update)
	api.Delete("/chat/delete/:fid", middleware.DeserializeUser, delete)

	api.Get("/chat", middleware.DeserializeUser, getAll)
	api.Get("/chat/:fid", middleware.DeserializeUser, getById)
	api.Get("/chat/user/:fid", middleware.DeserializeUser, getByUserId)
}

func create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)

	return nil
}

func update(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)

	return nil

}

func delete(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}

func getAll(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}

func getById(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}

func getByUserId(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}
