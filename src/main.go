package main

import (
	"github.com/gofiber/fiber/v2"
	router "github.com/notAlyosha/quiz-go/api/routes/quiz"
)

func main() {

	// Initialize a new Fiber app
	app := fiber.New()

	api := app.Group("/api")
	{
		router.SetupQuizRouter(app)
	}
	// Start the server on port 3000
	app.Listen(":3000")

}
