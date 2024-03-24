package user

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"

	"github.com/notAlyosha/quiz-go/pkg/middleware"
)

func SetupUserRouter(api fiber.Router) {
	api.Post("/user/add", middleware.DeserializeUser, create)
	api.Patch("/user/update/:fid", middleware.DeserializeUser, update)
	api.Delete("/user/delete/:fid", middleware.DeserializeUser, delete)

	api.Get("/user", middleware.DeserializeUser, geyByAuthData)
	api.Get("/user", middleware.DeserializeUser, getAll)
	api.Get("/user/:fid", middleware.DeserializeUser, getById)
	api.Get("/user/:fid/groups", middleware.DeserializeUser, getByGroupId)
	api.Get("/user/:fid/role", middleware.DeserializeUser, getByRoleId)

}

// create new user
func create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(entityUser.UserResponse)

	var newUser entityUser.UserCreate
	ctx.BodyParser(newUser)

	err := newUser.Check()

	if err != nil {
		errArray := strings.Split(err.Error(), ";")
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errArray)
	}

	// todo validate new user

	return createService(ctx, user, newUser)
}

func update(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")

	user := ctx.Locals("user").(entityUser.UserResponse)

	var newUser entityUser.UserCreate

	ctx.BodyParser(newUser)

	return updateService(ctx, user, newUser, fid)

}

func delete(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	user := ctx.Locals("user").(entityUser.UserResponse)

	ufid, err := uuid.FromString(fid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	return deleteService(ctx, user, ufid)

}

func geyByAuthData(ctx *fiber.Ctx) error {
	return nil
}

func getAll(ctx *fiber.Ctx) error {
	//user := ctx.Locals("user").(entityUser.UserResponse)
	return nil

}

func getById(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	// user := ctx.Locals("user").(entityUser.UserResponse)

	ufid, err := uuid.FromString(fid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	return getByIdService(ctx, ufid)

}

func getByGroupId(ctx *fiber.Ctx) error {
	fid := ctx.Params("fid")
	user := ctx.Locals("user").(entityUser.UserResponse)

	ufid, err := uuid.FromString(fid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	return getByGroupIdService(ctx, user, ufid)

}

func getByRoleId(ctx *fiber.Ctx) error {
	return getByRoleIdService("")

}
