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
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
	String()
}

func (l Lift) printLiftType() {
	fmt.Printf("sending to %v\n\n", []string{"small lift", "standard lift", "large lift"}[int(l)])
}

type Mototcycles string
type Cars string
type Trucks string

func (m Mototcycles) PickLift() Lift {
	return SmallLift
}

func (m Mototcycles) String() {
	fmt.Printf("--Vehicle: %v--\n", string(m))
}

func (c Cars) PickLift() Lift {
	return StandardLift
}

func (c Cars) String() {
	fmt.Printf("--Vehicle: %v--\n", string(c))
}

func (t Trucks) PickLift() Lift {
	return LargeLift
}

func (t Trucks) String() {
	fmt.Printf("--Vehicle: %v--\n", string(t))
}

func pickLifts(vehicles []LiftPicker) {
	fmt.Println("Picking Lifts for:")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		vehicle.String()
		lift := vehicle.PickLift()
		lift.printLiftType()
	}
	fmt.Println()
}

func main() {
	vehicles := []LiftPicker{Trucks("Road Devourer"), Mototcycles("Classic"), Cars("Porche"), Mototcycles("Sports"), Cars("Lamborghini")}

	pickLifts(vehicles)
}
