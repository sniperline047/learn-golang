package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(a, b int) int {
	return a + b
}

func greet() {
	fmt.Println("Hellooooo!!")
}

func main() {
	greet()

	var dozen int = double(6)
	fmt.Println("Dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("Baker's Dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another Baker's Dozen is", anotherBakersDozen)
}
