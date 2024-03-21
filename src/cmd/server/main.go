package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/internal/account"
	"github.com/notAlyosha/quiz-go/internal/config"
	"github.com/notAlyosha/quiz-go/internal/group"
	"github.com/notAlyosha/quiz-go/internal/order"
	"github.com/notAlyosha/quiz-go/internal/quiz"
	"github.com/notAlyosha/quiz-go/internal/session"
	"github.com/notAlyosha/quiz-go/internal/subject"
	"github.com/notAlyosha/quiz-go/internal/user"
)

func main() {

	config := config.LoadConfig()

	// TODO create database instance

	// create fiber app
	app := fiber.New()

	// mount routes
	api := app.Group("/api")
	account.SetupAccountRouter(api)
	user.SetupUserRouter(api)
	group.SetupGroupRouter(api)
	order.SetupOrderRoutes(api)
	quiz.SetupQuizRouter(api)
	session.SetupSessiontRoutes(api)
	subject.SetupSubjectRoute(api)

	// listen and serve
	app.Listen(":" + config.ServerPort)

}
