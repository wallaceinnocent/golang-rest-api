package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wallaceinnocent/go-rest-api/internals/controllers"
	"github.com/wallaceinnocent/go-rest-api/internals/services"
)

func CreateRoutes(app *fiber.App) {

	userServie := services.NewUserService()
	usersController := controllers.NewUsersController(userServie)

	app.Get("/", usersController.GetUser)

}
