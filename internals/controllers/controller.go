package controllers

import "github.com/gofiber/fiber/v3"

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (u *UsersController) GetUser(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"name": "User1"})
}
