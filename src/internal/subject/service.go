package subject

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/notAlyosha/quiz-go/internal/entity/subject"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"
)

func createService(ctx *fiber.Ctx, user entityUser.UserResponse, newSubject entity.SubjectInput) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot create subject"})
	}

	if user.Role == "Teacher" || user.Role == "Admin" {
		// build subject
		// Todo save new subject in database

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Subject has been successfully created"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func updateService(ctx *fiber.Ctx, user entityUser.UserResponse, newSubject entity.SubjectInput, subjectFid uuid.UUID) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot change subject"})
	}

	if user.Role == "Teacher" {
		// todo check that teacher has access to it's subject
		// todo change group in db
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Subject has been successfully updated"})

	}

	if user.Role == "Admin" {
		// Todo update subject in db
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Subject has been successfully updated"})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func deleteService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	if user.Role == "Student" || user.Role == "Teacher" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "You cannot delete this subject"})

	}

	if user.Role == "Admin" {

		// todo Get subject record from database and set data in subject
		var subject entity.Subject

		subject.IsDeleted = true
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Subject has been successfully deleted"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func getAllService(ctx *fiber.Ctx, user entityUser.UserResponse, fid string) error {
	var foundSubjects *[]entity.Subject

	if user.Role == "Student" || user.Role == "Teacher" {
		// todo check that user has access to subject with given fid
		// todo set record data to foundSubject structure
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(foundSubjects)

	}

	if user.Role == "Admin" {
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(foundSubjects)
	}

	// todo role check

	return ctx.Status(fiber.StatusFound).JSON(foundSubjects)

}

func getByIdService(ctx *fiber.Ctx, user entityUser.UserResponse, fid string) error {
	// variable that stores users in subject with fid
	// todo get data from db
	var subject entity.Subject

	if user.Role == "Student" || user.Role == "Teacher" {
		// todo check that user has access to subject with given fid
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(subject)

	}

	if user.Role == "Admin" {
		// todo get data from db

		return ctx.Status(fiber.StatusFound).JSON(subject)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}
