package quiz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupQuizRouter(api fiber.Router) {
	api.Post("/quiz/add", middleware.DeserializeUser, create)
	api.Patch("/quiz/update/:fid", middleware.DeserializeUser, update)
	api.Delete("/quiz/delete/:fid", middleware.DeserializeUser, delete)
	api.Get("/quiz/", middleware.DeserializeUser, getAll)
	api.Get("/quiz/:fid", middleware.DeserializeUser, getById)
	api.Get("/quiz/group/:fid", middleware.DeserializeUser, getByGroupId)
	api.Get("/quiz/user/:fid", middleware.DeserializeUser, getByUserId)

}

func create(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return ctx.SendString("It works! Yay")
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

func getByGroupId(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}

func getByUserId(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}
