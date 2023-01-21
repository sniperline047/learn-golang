//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func greetPerson(name string) {
	fmt.Println("Hola!", name)
}

func secretMessage() string {
	return "420 > 69"
}

func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func randomNumber() int {
	return rand.Int()
}

func twoRandomNumber() (int, int) {
	return rand.Int(), rand.Int()
}

func main() {
	greetPerson("sniperline047")

	fmt.Println("Sum of 3 number is: ", addThreeNumbers(69, 420, 911))

	fmt.Println("Secret message is:", secretMessage())

	fmt.Println("Random number is: ", randomNumber())
	rand1, ranr2 := twoRandomNumber()
	fmt.Println("Two Random number is: ", rand1, "and", ranr2)

	fmt.Println("Sum of 3 random number is: ", addThreeNumbers(randomNumber(), rand1, ranr2))
}
