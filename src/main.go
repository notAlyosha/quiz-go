package main

import (
	"github.com/gofiber/fiber/v2"
	router "github.com/notAlyosha/quiz-go/api/routes"
)

func main() {
	//Initialize a new Fiber app
	app := fiber.New()

	api := app.Group("/api")

	router.SetupAccountRouter(api)
	router.SetupUserRouter(api)
	router.SetupQuizRouter(api)
	router.SetupGroupRouter(api)
	router.SetupOrderRoutes(api)
	router.SetupChatRouter(api)

	// Start the server on port 3000
	app.Listen(":8080")

}
