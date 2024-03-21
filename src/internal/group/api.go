package group

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/notAlyosha/quiz-go/internal/entity/group"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"

	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupGroupRouter(api fiber.Router) {
	api.Post("/group/add", middleware.DeserializeUser, create)
	api.Patch("/group/update/:fid", middleware.DeserializeUser, update)
	api.Delete("/group/delete/:fid", middleware.DeserializeUser, delete)

	api.Get("/group", getAll)
	api.Get("/group/:fid", getById)
}

func create(ctx *fiber.Ctx) error {
	var group *entity.GroupInput
	ctx.BodyParser(&group)

	user := ctx.Locals("user").(entityUser.UserResponse)

	return createService(ctx, user, *group)
}

func update(ctx *fiber.Ctx) error {

	fid := ctx.Params("fid")
	ufid, err := uuid.FromString(fid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	var group *entity.GroupInput
	ctx.BodyParser(&group)

	user := ctx.Locals("user").(entityUser.UserResponse)
	return updateService(ctx, user, *group, ufid)
}

func delete(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	ufid, err := uuid.FromString(fid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	user := ctx.Locals("user").(entityUser.UserResponse)
	return deleteService(ctx, user, ufid)
}

func getAll(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	user := ctx.Locals("user").(entityUser.UserResponse)
	return getAllService(ctx, user, fid)
}

func getById(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")

	user := ctx.Locals("user").(entityUser.UserResponse)

	return getByIdService(ctx, user, fid)
}
