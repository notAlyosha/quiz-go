package session

import (
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"

	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupSessiontRoutes(api fiber.Router) {
	api.Post("/subject/add", middleware.DeserializeUser, add)
	api.Patch("/subject/update", middleware.DeserializeUser, update)
	api.Get("/subject", middleware.DeserializeUser, getAll)
	api.Get("/subject/:fid", middleware.DeserializeUser, getByID)
	api.Get("/subject/user/:fid", middleware.DeserializeUser, getByUserID)
	api.Get("/subject/group/:fid", middleware.DeserializeUser, getByGroupID)
}

func add(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	return nil
}

func update(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil
}

func getAll(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil
}

func getByID(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil
}

func getByUserID(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil
}

func getByGroupID(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil
}
