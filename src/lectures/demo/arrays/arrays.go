package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkedCleanedRooms(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	// var rooms [4]Room

	// rooms[0] = Room{"1", true}
	// rooms[1] = Room{"2", true}
	// rooms[2] = Room{"3", true}
	// rooms[3] = Room{name: "4"}

	rooms := [...]Room{
		{name: "Room 1"},
		{name: "Room 2"},
		{name: "Room 3", cleaned: true},
		{name: "Room 4", cleaned: true},
	}

	fmt.Println("Performing cleaning checks...")
	checkedCleanedRooms(rooms)

	rooms[3].cleaned = false

	fmt.Println("Re-Performing cleaning checks...")
	checkedCleanedRooms(rooms)
}
