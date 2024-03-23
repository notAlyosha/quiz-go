package subject

import (
	entity "github.com/notAlyosha/quiz-go/internal/entity/subject"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupSubjectRoute(api fiber.Router) {
	api.Post("/subject/add", middleware.DeserializeUser, add)
	api.Patch("/subject/update/:fid", middleware.DeserializeUser, update)
	api.Patch("/subject/delete/:fid", middleware.DeserializeUser, delete)
	api.Get("/subject/", middleware.DeserializeUser, getAll)
	api.Get("/subject/:fid", middleware.DeserializeUser, getByID)
	api.Get("/subject/user/:fid", middleware.DeserializeUser, getByUserID)
	api.Get("/subject/group/:fid", middleware.DeserializeUser, getByGroupID)
}

func add(ctx *fiber.Ctx) error {
	subject := &entity.SubjectInput{}

	ctx.BodyParser(&subject)
	user := ctx.Locals("user").(entityUser.UserResponse)
	return createService(ctx, user, *subject)
}

func update(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")

	ufid, err := uuid.FromString(fid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	subject := &entity.SubjectInput{}

	ctx.BodyParser(&subject)

	user := ctx.Locals("user").(entityUser.UserResponse)
	return updateService(ctx, user, *subject, ufid)
}

func delete(ctx *fiber.Ctx) error {
	fid := uuid.FromStringOrNil(ctx.Params("fid"))
	user := ctx.Locals("user").(entityUser.UserResponse)
	return deleteService(ctx, user, fid)
}

func getAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)
	return getAllService(ctx, user)
}

func getByID(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	user := ctx.Locals("user").(entityUser.UserResponse)
	return getByIdService(ctx, user, fid)
}

func getByUserID(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	user := ctx.Locals("user").(entityUser.UserResponse)
	return getByUserIDService(ctx, user, fid)
}

func getByGroupID(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	user := ctx.Locals("user").(entityUser.UserResponse)
	return geyByGroupIDService(ctx, user, fid)
}
