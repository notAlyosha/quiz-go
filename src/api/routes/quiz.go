package router

import (
	"github.com/gofiber/fiber/v2"
	quizHandler "github.com/notAlyosha/quiz-go/api/controllers/quiz"
)

func SetupQuizRouter(api fiber.Router) {
	api.Get("/quiz/", quizHandler.GetAll)
	api.Get("/quiz/:id", quizHandler.GetById)
	api.Get("/quiz/group/:id", quizHandler.GetByGroupId)
	api.Get("/quiz/user/:id", quizHandler.GetByUserId)
	api.Get("/quiz/role/:id", quizHandler.GetByRoleId)

	api.Post("/quiz/add", quizHandler.Create)
	api.Post("/quiz/update/:id", quizHandler.Update)
	api.Post("/quiz/delete/:id", quizHandler.Delete)

}
