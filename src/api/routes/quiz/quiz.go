package router

import (
	"github.com/gofiber/fiber/v2"
	quizHandler "github.com/notAlyosha/quiz-go/api/controllers/quiz"
)

func SetupQuizRouter(app *fiber.App) {
	app.Get("/quiz/", quizHandler.GetAll)
	app.Get("/quiz/:id", quizHandler.GetById)
	app.Get("/quiz/group/:id", quizHandler.GetByGroupId)
	app.Get("/quiz/user/:id", quizHandler.GetByUserId)
	app.Get("/quiz/role/:id", quizHandler.GetByRoleId)

	app.Post("/quiz/add", quizHandler.Create)
	app.Post("/quiz/update/:id", quizHandler.Update)
	app.Post("/quiz/delete/:id", quizHandler.Delete)

}
