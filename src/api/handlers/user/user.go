package userHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/api/types/userType"
)

func Create(ctx *fiber.Ctx) error {
	var u *userType.User

	err := ctx.BodyParser(&u)

	if err != nil {
		return ctx.JSON(err)
	}

	if u.IsDeleted == nil {

		return ctx.JSON("IsDeleted is null")
	}

	ctx.SendStatus(fiber.StatusAccepted)
	return ctx.JSON(u)
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

func GetGroupById(ctx *fiber.Ctx) error {
	return nil

}

func GetByRoleId(ctx *fiber.Ctx) error {
	return nil

}
