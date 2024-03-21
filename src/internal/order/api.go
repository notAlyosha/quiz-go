package order

import (
	"github.com/gofiber/fiber/v2"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupOrderRoutes(api fiber.Router) {
	api.Post("order/add", middleware.DeserializeUser, create)
	api.Patch("/order/update/:id", middleware.DeserializeUser, update)
	api.Get("/order", middleware.DeserializeUser, getAll)
	api.Get("/order/:fid", middleware.DeserializeUser, getByID)
	api.Get("/order/user/:fid", middleware.DeserializeUser, getByUserID)
	api.Get("/order/group/:fid", middleware.DeserializeUser, getByGroupID)
}

func create(ctx *fiber.Ctx) error {
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
