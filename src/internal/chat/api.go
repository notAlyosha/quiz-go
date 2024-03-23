package chat

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"

	entity "github.com/notAlyosha/quiz-go/internal/entity/chat"
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

	var newChat entity.ChatInput
	ctx.BodyParser(newChat)

	return createService(ctx, user, newChat)
}

func update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)

	var newChat entity.ChatInput
	ctx.BodyParser(newChat)

	return updateService(ctx, user, newChat)

}

func delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	fid, _ := uuid.FromString(ctx.Get("fid"))
	return deleteService(ctx, user, fid)

}

func getAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	return getAllService(ctx, user)

}

func getById(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	fid, _ := uuid.FromString(ctx.Get("fid"))

	var newChat entity.ChatInput
	ctx.BodyParser(newChat)

	return getByIdService(ctx, user, newChat, fid)

}

func getByUserId(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)

	var newChat entity.ChatInput
	ctx.BodyParser(newChat)

	fid, _ := uuid.FromString(ctx.Get("fid"))

	return getByUserIdService(ctx, user, newChat, fid)

}
