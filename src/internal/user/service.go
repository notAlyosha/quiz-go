package user

import (
	"github.com/gofiber/fiber/v2"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/pkg/middleware"
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

		count, err := countProfileInDB(newUser)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		if count != 0 {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "User already exists"})
		}

		salt, err := middleware.GenerateRandomSalt(255)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		newUser.Password = middleware.HashPassword(newUser.Password, salt)
		FID := uuid.NewV4().String()

		// TODO TEST BYTE[255] => STRING CONVERTION IN MySQL DBMS
		if createUserInStorage(FID, string(salt), newUser) != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully created"})
	}

	if user.Role == "Admin" {

		count, err := countProfileInDB(newUser)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		if count != 0 {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "User already exists"})
		}

		salt, err := middleware.GenerateRandomSalt(255)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		newUser.Password = middleware.HashPassword(newUser.Password, salt)
		FID := uuid.NewV4().String()
		// TODO TEST BYTE[255] => STRING CONVERTION IN MySQL DBMS
		if createUserInStorage(FID, string(salt), newUser) != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully created"})
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})
}

func updateService(ctx *fiber.Ctx, user entityUser.UserResponse, newUser entityUser.UserCreate, updated_user_FID string) error {
	if user.Role == "Student" {
		if newUser.Role != "Student" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot change self role"})
		}

		count, err := countProfileInDB(newUser)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		if count != 0 {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "User already exists"})
		}

		if isUserInDB(updated_user_FID) != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "User not found"})
		}

		salt, err := getSaltUser(updated_user_FID)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		newUser.Password = middleware.HashPassword(newUser.Password, []byte(salt))

		if updateUserInStorage(updated_user_FID, newUser) != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})
	}

	if user.Role == "Teacher" {
		if newUser.Role != "Teacher" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Teacher cannot change role"})
		}

		count, err := countProfileInDB(newUser)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		if count != 0 {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Phone or Email or login already exists"})
		}

		if isUserInDB(updated_user_FID) != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "User not found"})
		}

		salt, err := getSaltUser(updated_user_FID)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		newUser.Password = middleware.HashPassword(newUser.Password, []byte(salt))

		if updateUserInStorage(updated_user_FID, newUser) != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})

	}

	if user.Role == "Admin" {
		count, err := countProfileInDB(newUser)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		if count != 0 {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Phone or Email or login already exists"})
		}

		if isUserInDB(updated_user_FID) != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "User not found"})
		}

		salt, err := getSaltUser(updated_user_FID)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		newUser.Password = middleware.HashPassword(newUser.Password, []byte(salt))

		if updateUserInStorage(updated_user_FID, newUser) != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully updated"})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func deleteService(ctx *fiber.Ctx, user entityUser.UserResponse, deleted_user_fid uuid.UUID) error {
	if user.Role == "Student" || user.Role == "Teacher" {
		if user.FrontID != deleted_user_fid {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Your id and provided id is not the same"})
		}

		if isUserInDB(deleted_user_fid.String()) == nil {
			err := deleteUserFromStorage(deleted_user_fid.String())
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
			}
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully deleted"})
	}

	if user.Role == "Admin" {

		if isUserInDB(deleted_user_fid.String()) == nil {
			err := deleteUserFromStorage(deleted_user_fid.String())
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unhandled server error"})
			}
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User has been successfully deleted"})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})
}

func getByIdService(ctx *fiber.Ctx, fid uuid.UUID) error {
	foundUser, err := userByIdFromStorage(fid)

	if err == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Provided id is not exsist"})
	}

	clientUser := entityUser.UserResponse{}

	clientUser.FrontID = uuid.FromStringOrNil(foundUser.FrontID)
	clientUser.Role = foundUser.Role

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
func getByRoleIdService() error {
	return nil
}
