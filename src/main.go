package main

import (
	"github.com/gofiber/fiber/v2"
	router "github.com/notAlyosha/quiz-go/api/routes"
)

func main() {

	// Initialize a new Fiber app
	app := fiber.New()

	api := app.Group("/api")

	router.SetupQuizRouter(api)
	router.SetupChatRouter(api)
	router.SetupUserRouter(api)
	router.SetupGroupRouter(api)

	// Start the server on port 3000
	app.Listen(":8080")

}
