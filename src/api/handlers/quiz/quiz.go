package quizHandler

import "github.com/gofiber/fiber/v2"

func Create(ctx *fiber.Ctx) error {
	return ctx.SendString("It works! Yay")
}

func Update(ctx *fiber.Ctx) error {
	return nil

}

func Delete(ctx *fiber.Ctx) error {
	return nil

}

func GetAll(ctx *fiber.Ctx) error {
	return nil

}

func GetById(ctx *fiber.Ctx) error {
	return nil

}

func GetByGroupId(ctx *fiber.Ctx) error {
	return nil

}

func GetByUserId(ctx *fiber.Ctx) error {
	return nil

}
