package main

import (
	"fmt"
	"main/internals/models"
)

func main() {
	name := another()
	fmt.Println(name)

}

func another() string {

	vehicle := models.Vehicle{RegNo: 1, Name: "crown", Model: "toyota"}
	user := &models.Users{}

	user.Age = 10
	user.Name = "User"
	user.VehicleReg = vehicle.RegNo

	changeStruct(user)

	return user.Name
}

func changeStruct(user *models.Users) {
	user.Name = "User2"
}
