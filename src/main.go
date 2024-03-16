package main

import (
	"crypto/rand"
	"crypto/rsa"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	router "github.com/notAlyosha/quiz-go/api/routes"
)

func main() {
	//Initialize a new Fiber app
	app := fiber.New()

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
	}))

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
