package main

import "fmt"

func main() {
	slice := []string{"No", "I", "Am", "Your", "Father", "!"}

	for i, element := range slice {
		fmt.Println(i, element)

		// range can be used to access each rune, even though they maybe of multiple bits (special chars)
		for _, ch := range element {
			fmt.Printf("%q\n", ch)
		}

		fmt.Println()
	}
}
