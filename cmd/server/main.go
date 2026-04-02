package main

import (
	"log"
	"main/internals/routes"

	"github.com/gofiber/fiber/v3"
)

func init() {
	log.Println("Starting ....")
}
func main() {
	app := fiber.New()

	routes.CreateRoutes(app)
	app.Listen(":8080")

}
