package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNumber int) {
	if spaceNumber > len(lot.spaces) || lot.spaces[spaceNumber-1].occupied {
		fmt.Println("Cannot park here, Occupied!")
		return
	}

	lot.spaces[spaceNumber-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNumber int) {
	if spaceNumber > len(lot.spaces) || lot.spaces[spaceNumber-1].occupied {
		fmt.Println("Cannot park here, Occupied!")
		return
	}

	lot.spaces[spaceNumber-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNumber int) {
	if spaceNumber > len(lot.spaces) || !lot.spaces[spaceNumber-1].occupied {
		fmt.Println("Space is already vacated or does not exist!")
		return
	}

	lot.spaces[spaceNumber-1].occupied = false
}

func main() {
	parkingLot := ParkingLot{make([]Space, 4)}
	fmt.Println("Initial:", parkingLot)

	occupySpace(&parkingLot, 2)
	fmt.Println("Updated:", parkingLot)

	parkingLot.occupySpace(3)
	fmt.Println("Updated:", parkingLot)

	parkingLot.vacateSpace(2)
	fmt.Println("Updated:", parkingLot)

	parkingLot.occupySpace(5)
}
