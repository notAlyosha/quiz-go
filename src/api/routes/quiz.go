package router

import (
	"github.com/gofiber/fiber/v2"
	quizHandler "github.com/notAlyosha/quiz-go/api/controllers/quiz"
)

func SetupQuizRouter(api fiber.Router) {
	api.Get("/quiz/", quizHandler.GetAll)
	api.Get("/quiz/:fid", quizHandler.GetById)
	api.Get("/quiz/group/:fid", quizHandler.GetByGroupId)
	api.Get("/quiz/user/:fid", quizHandler.GetByUserId)

	api.Post("/quiz/add", quizHandler.Create)
	api.Patch("/quiz/update/:fid", quizHandler.Update)
	api.Delete("/quiz/delete/:fid", quizHandler.Delete)

}
