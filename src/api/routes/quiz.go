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
	api.Get("/quiz/role/:fid", quizHandler.GetByRoleId)

	api.Post("/quiz/add", quizHandler.Create)
	api.Post("/quiz/update/:fid", quizHandler.Update)
	api.Post("/quiz/delete/:fid", quizHandler.Delete)

}
