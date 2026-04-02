package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func init() {
	log.Fatal("Startes")
}
func main() {
	app := fiber.New()
	app.Listen(":8080")

}
