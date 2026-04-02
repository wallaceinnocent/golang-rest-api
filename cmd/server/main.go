package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/wallaceinnocent/go-rest-api/internals/routes"
)

func init() {
	log.Println("Starting ....")
}
func main() {
	app := fiber.New()

	routes.CreateRoutes(app)
	app.Listen(":8080")

}
