package order

import (
	"github.com/gofiber/fiber/v2"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"
)

func createService(ctx *fiber.Ctx, user entityUser.UserResponse, newUser entityUser.UserCreate) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot create users"})
	}

	if user.Role == "Teacher" {
		if newUser.Role == "Admin" || newUser.Role == "Teacher" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Teacher cannot create admin or teacher"})
		}

		// Todo save new user in database:=

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully created"})

	}

	if user.Role == "Admin" {
		// Todo save new user in database

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully created"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func updateService(ctx *fiber.Ctx, user entityUser.UserResponse, newUser entityUser.UserCreate) error {
	if user.Role == "Student" {
		if newUser.Role != "Student" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot change role"})
		}

		// Todo update user in db

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})

	}

	if user.Role == "Teacher" {
		if newUser.Role != "Teacher" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Teacher cannot change role"})
		}

		// Todo update user in db

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})

	}

	if user.Role == "Admin" {
		// Todo update user in db

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func deleteService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	if user.Role == "Student" || user.Role == "Teacher" {
		if user.FrontID != fid {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Your id and provided id is not the same"})
		}

		// todo Get user record from database and set data in user
		//var user entityUser.User

		//user.IsDeleted = true
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully deleted"})
	}

	if user.Role == "Admin" {

		// todo Get user record from database and set data in user
		//var user entityUser.User

		//user.IsDeleted = true
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully deleted"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func getAllService(ctx *fiber.Ctx, user entityUser.UserResponse) error {
	return nil
}

func getByIDService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}

func getByUserIDService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}
