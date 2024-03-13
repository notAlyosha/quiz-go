package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// config := database.Config{}

	// secureDB, err := database.InitDB(config)

	// mainDB, err := database.InitDB(config)

	// Initialize a new Fiber app
	app := fiber.New()

	// Start the server on port 3000
	app.Listen(":3000")

}
