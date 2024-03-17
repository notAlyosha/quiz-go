package main

import (
	"github.com/gofiber/fiber/v2"
	router "github.com/notAlyosha/quiz-go/api/routes"
)

func main() {
	//Initialize a new Fiber app
	app := fiber.New()

	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	panic(err)
	// }

	api := app.Group("/api")

	// unauthentificated route
	router.SetupAccountRouter(api)

	// JWT Middleware
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{
	// 		JWTAlg: jwtware.RS256,
	// 		Key:    privateKey.Public(),
	// 	},
	// }))

	// app.Use(auth.JwtMiddleware)

	// authentification required
	router.SetupUserRouter(api)
	router.SetupQuizRouter(api)
	router.SetupGroupRouter(api)
	router.SetupOrderRoutes(api)
	router.SetupChatRouter(api)

	// Start the server on port 3000
	app.Listen(":8080")

}
