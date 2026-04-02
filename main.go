package main

import (
	"fmt"
	"main/models"
)

func main() {
	name, reg, model := another()
	fmt.Println()
	fmt.Println("Vehicle name:", name)
	fmt.Println()
	fmt.Printf("The reg is %d and model is %s", reg, model)
	fmt.Println()

}

func another() (string, int, string) {
	vehicle := models.Vehicle{RegNo: 1, Name: "crown", Model: "toyota"}

	return vehicle.Name, vehicle.RegNo, vehicle.Model
}
