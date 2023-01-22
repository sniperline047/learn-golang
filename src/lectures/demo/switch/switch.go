package main

import (
	"fmt"
)

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheat price")
	case p < 10:
		fmt.Println("Moderately priced item")
	case p > 10:
		fmt.Println("Expensive Item")
	default:
		fmt.Println("Invalid Item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy Ticket")
	case Business:
		fmt.Println("Business Ticket")
	case FirstClass:
		fmt.Println("First Class Ticket")
	default:
		fmt.Println("Invalid Ticket")
	}
}
