package account

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/notAlyosha/quiz-go/internal/config"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func SetupAccountRouter(api fiber.Router) {
	api.Post("/enter/signIn", signIn)
	api.Post("/enter/signUp", signUp)
	api.Post("/enter/refresh", refreshAccessToken)
	api.Post("/enter/logout", logoutUser)
}

func signUp(ctx *fiber.Ctx) error {
	var payload *entityUser.SignInInput

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Todo: implement validation

	// Todo: implement password confirm

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	newUser := entityUser.User{}

	return ctx.JSON(newUser)
}

func signIn(ctx *fiber.Ctx) error {
	var payload *entityUser.SignInInput

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Todo implement validation

	// Todo find user in database and paste record in structure below
	var user entityUser.User

	// Todo check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Salt), []byte(payload.Password))
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "Password is not valid"})
	}

	config := config.LoadConfig()

	accessTokenDetails, err := utils.CreateToken(user.FrontID.String(), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(user.FrontID.String(), config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// set cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.ServerHost,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.ServerHost,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   config.ServerHost,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

func refreshAccessToken(ctx *fiber.Ctx) error {

	refresh_token := ctx.Cookies("refresh_token")

	if refresh_token == "" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "could not refresh access token"})
	}

	config := config.LoadConfig()

	// tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	// }

	var user entityUser.User
	// find user in database using token claims and paste a record into structure above

	accessTokenDetails, err := utils.CreateToken(user.FrontID.String(), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.ServerHost,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   config.ServerHost,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

func logoutUser(ctx *fiber.Ctx) error {
	refresh_token := ctx.Cookies("refresh_token")

	if refresh_token == "" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token is invalid or session has expired"})
	}

	config := config.LoadConfig()

	_, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	expired := time.Now().Add(-time.Hour * 24)
	ctx.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
