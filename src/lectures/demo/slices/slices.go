package main

import "fmt"

func printSlice(title string, routes []string) {
	fmt.Println()
	fmt.Println("==========", title, "===================")

	for i := 0; i < len(routes); i++ {
		route := routes[i]
		fmt.Println(route)
	}
}

func main() {
	routes := []string{"Tashi Station", "Blue Milk Store", "Jaba's place"}
	printSlice("Route 1", routes)

	fmt.Println()
	fmt.Println("Updating routes...")
	routes = append(routes, "Worm Pit")
	printSlice("Route 2", routes)

	fmt.Println()
	fmt.Println(routes[0], "Visited!")
	fmt.Println(routes[1], "Visited!")

	routes = routes[2:]

	printSlice("Remaining routes", routes)
}
