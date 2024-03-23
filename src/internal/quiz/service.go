package quiz

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/notAlyosha/quiz-go/internal/entity/quiz"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"
)

func createService(ctx *fiber.Ctx, user entityUser.UserResponse, newQuiz entity.QuizInput) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot create quiz"})
	}

	if user.Role == "Teacher" || user.Role == "Admin" {
		// build quiz
		// Todo save new quiz in database

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Quiz has been successfully created"})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func updateService(ctx *fiber.Ctx, user entityUser.UserResponse, newQuiz entity.QuizInput, subjectFid uuid.UUID) error {
	if user.Role == "Student" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Student cannot change quiz"})
	}

	if user.Role == "Teacher" {

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Quiz has been successfully updated"})

	}

	if user.Role == "Admin" {
		// Todo update quiz in db
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Quiz has been successfully updated"})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid role"})

}

func deleteService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}

func getAllService(ctx *fiber.Ctx, user entityUser.UserResponse) error {
	return nil
}

func getByIdService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}

func getByGroupIDService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}

func getUserByIdService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}
