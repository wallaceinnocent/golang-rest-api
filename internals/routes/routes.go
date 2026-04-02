package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wallaceinnocent/go-rest-api/internals/controllers"
)

func CreateRoutes(app *fiber.App) {
	usersController := controllers.NewUsersController()

	app.Get("/", usersController.GetUser)

}
