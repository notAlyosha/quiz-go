package user

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

		createUser := convertFromUC2U(newUser)

		// Todo save new user in database

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully created"})

	}

	if user.Role == "Admin" {
		createUser := convertFromUC2U(newUser)

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

		createUser := convertFromUC2U(newUser)

		// Todo update user in db

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})

	}

	if user.Role == "Teacher" {
		if newUser.Role != "Teacher" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Teacher cannot change role"})
		}

		createUser := convertFromUC2U(newUser)

		// Todo update user in db

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})

	}

	if user.Role == "Admin" {
		createUser := convertFromUC2U(newUser)

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
		var user entityUser.User

		user.IsDeleted = true
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully deleted"})
	}

	if user.Role == "Admin" {

		// todo Get user record from database and set data in user
		var user entityUser.User

		user.IsDeleted = true
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully deleted"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func getByIdService(ctx *fiber.Ctx, fid uuid.UUID) error {
	var foundUser *entityUser.User
	// todo set record data to user structure

	if foundUser == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Provided id is not exsist"})
	}

	clientUser := convertFromU2UC(*foundUser)

	return ctx.Status(fiber.StatusFound).JSON(clientUser)

}

func getGroupByIdService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	// variable that stores users in group with fid
	// todo get data from db
	var users []entityUser.User

	if user.Role == "Student" || user.Role == "Teacher" {
		// todo check that user has access to group with given fid

		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(users)

	}

	if user.Role == "Admin" {
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(users)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func convertFromUC2U(newUser entityUser.UserCreate) entityUser.User {
	user := &entityUser.User{}

	user.Name = newUser.Name
	user.Login = newUser.Login
	user.Role = &newUser.Role
	user.Email = newUser.Email
	user.LogoURL = newUser.LogoURL

	return *user
}

func convertFromU2UC(user entityUser.User) entityUser.UserCreate {
	userCreate := &entityUser.UserCreate{}

	userCreate.Email = user.Email
	userCreate.Login = user.Login
	userCreate.LogoURL = user.LogoURL
	userCreate.Name = user.Name
	userCreate.Role = *user.Role

	return *userCreate
}
