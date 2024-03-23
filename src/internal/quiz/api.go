package quiz

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/notAlyosha/quiz-go/internal/entity/quiz"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/pkg/middleware"
	uuid "github.com/satori/go.uuid"
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
	user := ctx.Locals("user").(entityUser.UserResponse)
	var newQuiz entity.QuizInput
	ctx.BodyParser(&user)
	return createService(ctx, user, newQuiz)
}

func update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	var newQuiz entity.QuizInput
	ctx.BodyParser(&user)
	fid, _ := uuid.FromString(ctx.Get("fid"))
	return updateService(ctx, user, newQuiz, fid)

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
	return getByIdService(ctx, user, fid)

}

func getByGroupId(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	fid, _ := uuid.FromString(ctx.Get("fid"))
	return getByGroupIDService(ctx, user, fid)

}

func getByUserId(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	fid, _ := uuid.FromString(ctx.Get("fid"))
	return getByGroupIDService(ctx, user, fid)

}
