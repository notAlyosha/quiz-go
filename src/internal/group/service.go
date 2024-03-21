package group

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/notAlyosha/quiz-go/internal/entity/group"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"
)

func createService(ctx *fiber.Ctx, user entityUser.UserResponse, newGroup entity.GroupInput) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot create groups"})
	}

	if user.Role == "Teacher" || user.Role == "Admin" {
		// build group

		// Todo save new group in database

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Group has been successfully created"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func updateService(ctx *fiber.Ctx, user entityUser.UserResponse, newGroup entity.GroupInput, groupFid uuid.UUID) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot change group"})
	}

	if user.Role == "Teacher" {
		// todo check that teacher has access to it's group
		// todo change group in db
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Group has been successfully updated"})

	}

	if user.Role == "Admin" {
		// Todo update group in db
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Group has been successfully updated"})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func deleteService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	if user.Role == "Student" || user.Role == "Teacher" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "You cannot delete this group"})

	}

	if user.Role == "Admin" {

		// todo Get group record from database and set data in group
		var group entity.Group

		group.IsDeleted = true
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Group has been successfully deleted"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func getAllService(ctx *fiber.Ctx, user entityUser.UserResponse, fid string) error {
	var foundGroups *[]entity.Group

	if user.Role == "Student" || user.Role == "Teacher" {
		// todo check that user has access to group with given fid
		// todo set record data to foundGroup structure
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(foundGroups)

	}

	if user.Role == "Admin" {
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(foundGroups)
	}

	// todo role check

	return ctx.Status(fiber.StatusFound).JSON(foundGroups)

}

func getByIdService(ctx *fiber.Ctx, user entityUser.UserResponse, fid string) error {
	// variable that stores users in group with fid
	// todo get data from db
	var group entity.Group

	if user.Role == "Student" || user.Role == "Teacher" {
		// todo check that user has access to group with given fid
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(group)

	}

	if user.Role == "Admin" {
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(group)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}
