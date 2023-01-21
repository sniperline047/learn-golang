package main

import "fmt"

func main() {
	// New variable
	var myName = "sniperline047"
	fmt.Println("My name is: ", myName)

	// New variable with type string
	var name string = "Sniper Line 047"
	fmt.Println("My name is: ", name, "is of type string")

	// intialise and assign
	usernmae := "snipe047"
	fmt.Println(usernmae)

	// New variable initialise and assignment with type int
	var sum int = 69
	fmt.Println("Sum is: ", sum, "obviosly!")

	// New variable with type int, defaults to 0 in uninitialsed
	var sum1 int
	fmt.Println("Sum1 is: ", sum1, "obviosly!")

	// init and assign multiple variables
	part1, part2 := 1, 5
	fmt.Println("Initialize and assign two parts, part1:", part1, "and part2 is: ", part2)

	// okay error idiom
	value1, errorVal := 1, false
	fmt.Println("Value1 is: ", value1, "and error is:", errorVal)
	// reuses the second variable errorVal as it was already created
	value2, errorVal := 0, true
	fmt.Println("Value1 is: ", value2, "and error is:", errorVal)

	// reassignemnt
	sum = part1 + part2
	fmt.Println("New sum is: ", sum)

	// declaration block
	var (
		lessonType1 string = "This is a string"
		lessonType2 int    = 420
	)
	fmt.Println(lessonType1, lessonType2)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)

	// both valk1 and valk2 are of int type
	var valk1, valk2 int = 1, 2
	fmt.Println(valk1, valk2)
}
