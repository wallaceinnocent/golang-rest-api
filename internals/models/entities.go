package models

type Users struct {
	Name string
	Age  int
	Vehicle
	VehicleReg int
}

type Vehicle struct {
	RegNo int
	Name  string
	Model string
}
