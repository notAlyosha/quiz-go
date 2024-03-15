package router

import (
	"github.com/gofiber/fiber/v2"
	orderHandler "github.com/notAlyosha/quiz-go/api/controllers/order"
)

func SetupOrderRoutes(api fiber.Router) {
	api.Post("order/add", orderHandler.Create)
	api.Patch("/order/update", orderHandler.Update)
	api.Get("/order", orderHandler.GetAll)
	api.Get("/order/:fid", orderHandler.GetByID)
	api.Get("/order/user/:fid", orderHandler.GetByUserID)
	api.Get("/order/group/:fid", orderHandler.GetByGroupID)
}
