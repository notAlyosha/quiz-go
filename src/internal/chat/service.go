package chat

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/notAlyosha/quiz-go/internal/entity/chat"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	uuid "github.com/satori/go.uuid"
)

func createService(ctx *fiber.Ctx, user entityUser.UserResponse, newChat entity.ChatInput) error {
	return nil
}

func updateService(ctx *fiber.Ctx, user entityUser.UserResponse, newChat entity.ChatInput) error {
	return nil
}

func deleteService(ctx *fiber.Ctx, user entityUser.UserResponse, fid uuid.UUID) error {
	return nil
}

func getAllService(ctx *fiber.Ctx, user entityUser.UserResponse) error {
	return nil
}

func getByIdService(ctx *fiber.Ctx, user entityUser.UserResponse, newChat entity.ChatInput, fid uuid.UUID) error {
	return nil
}

func getByUserIdService(ctx *fiber.Ctx, user entityUser.UserResponse, newChat entity.ChatInput, fid uuid.UUID) error {
	return nil
}
