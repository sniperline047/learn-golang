package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type MeleniumFalcon struct {
	Pilot   Passenger
	CoPilot Passenger
}

func main() {
	leia := Passenger{"Leia", 1, false}
	fmt.Println(leia)

	var (
		han   = Passenger{"Han", 2, false}
		chewy = Passenger{Name: "Chewbacca", TicketNumber: 3}
	)

	fmt.Println(han, chewy)

	var luke Passenger
	luke.Name = "Luke Skywalker"
	luke.Boarded = true
	luke.TicketNumber = 4

	fmt.Println(luke)

	falcon := MeleniumFalcon{han, chewy}

	han.Boarded = true
	chewy.Boarded = true

	if han.Boarded && chewy.Boarded {
		fmt.Println(falcon.Pilot.Name, "is flying the Melinium falcon and", falcon.CoPilot.Name, "is the CoPilot")
	} else {
		fmt.Println("Mission failed, Empire strikes back!")
	}
}
