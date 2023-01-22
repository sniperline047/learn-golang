package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum is", sum)

	// simple for loop
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	// for loop (while type implementation)
	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement, Sum is", sum)
	}

	// infinite loop
	for {
		if sum == 10 {
			fmt.Println("Stoping loop execution...")
			break
		}

		fmt.Println("Looping infinitely....")
	}
}
