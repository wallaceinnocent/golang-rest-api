package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wallaceinnocent/go-rest-api/internals/services"
)

type UsersController struct {
	Userservice *services.UserService
}

func NewUsersController(s *services.UserService) *UsersController {
	return &UsersController{
		Userservice: s,
	}
}

func (u *UsersController) GetUser(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"name": u.Userservice.GetUserData()})
}
