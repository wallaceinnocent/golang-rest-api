package routes

import (
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestRouteSetup(t *testing.T) {
	app := fiber.New()

	// Test if the app instance is created correctly
	assert.NotNil(t, app)
}
