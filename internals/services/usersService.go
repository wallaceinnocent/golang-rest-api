package services

import "github.com/wallaceinnocent/go-rest-api/internals/models"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserData() string {

	user := models.Users{Name: "User10"}
	return user.Name
}
