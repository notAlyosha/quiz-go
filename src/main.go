package main

import (
	"github.com/gofiber/fiber/v2"
)

var Server *fiber.App

func main() {

	// Initialize a new Fiber app
	Server = fiber.New()

	// Start the server on port 3000
	Server.Listen(":3000")

}
