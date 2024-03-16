package router

import (
	"github.com/gofiber/fiber/v2"
	subjectHandler "github.com/notAlyosha/quiz-go/api/handlers/subject"
)

func SetupSubjectRoute(api fiber.Router) {
	api.Post("/subject/add", subjectHandler.Add)
	api.Patch("/subject/update/:fid", subjectHandler.Update)
	api.Get("/subject/", subjectHandler.GetAll)
	api.Get("/subject/:fid", subjectHandler.GetByID)
	api.Get("/subject/user/:fid", subjectHandler.GetByUserID)
	api.Get("/subject/group/:fid", subjectHandler.GetByGroupID)
}
