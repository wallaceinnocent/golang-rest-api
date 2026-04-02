package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func init() {
	log.Println("Starting ....")
}
func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello world",
		})
	})
	app.Listen(":8080")

}
