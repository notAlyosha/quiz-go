package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/internal/account"
	"github.com/notAlyosha/quiz-go/internal/config"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/internal/group"
	"github.com/notAlyosha/quiz-go/internal/order"
	"github.com/notAlyosha/quiz-go/internal/quiz"
	"github.com/notAlyosha/quiz-go/internal/session"
	"github.com/notAlyosha/quiz-go/internal/subject"
	"github.com/notAlyosha/quiz-go/internal/user"
	"github.com/notAlyosha/quiz-go/pkg/dbcontext"
)

func main() {

	config := config.LoadConfig()

	// TODO create database instance
	err := dbcontext.CreateDBConnection()

	if err != nil {
		panic(err)
	}

	db := dbcontext.GetDB()

	q := db.NewQuery("SELECT * FROM users")

	var users *entityUser.User

	err = q.One(&users)
	if err != nil {
		fmt.Println(err)
	}

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
