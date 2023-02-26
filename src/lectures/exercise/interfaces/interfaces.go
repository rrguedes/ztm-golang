//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SMALL = iota
	STANDARD
	LARGE
)

type Lifter interface {
	liftVehicle()
}

type Vehicle struct {
	model       string
	vehicleType string
	liftType    int
}

type Motorcycle Vehicle
type Car Vehicle
type Truck Vehicle

func (v Motorcycle) liftVehicle() {
	fmt.Println("Lifting Motorcycle")
}

func (v Car) liftVehicle() {
	fmt.Println("Lifting Car")
}

func (v Truck) liftVehicle() {
	fmt.Println("Lifting Truck")
}

func liftVehicles(vehicles []Lifter) {
	fmt.Println("-- Initiating Service --")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		vehicle.liftVehicle()
	}
}

func main() {
	vehicles := []Lifter{
		Motorcycle{model: "Titan", vehicleType: "Road Blaster", liftType: SMALL},
		Car{model: "ASX", vehicleType: "SUV", liftType: STANDARD},
		Truck{model: "Globetrotter", vehicleType: "Bitruck", liftType: LARGE},
	}
	liftVehicles(vehicles)
}
