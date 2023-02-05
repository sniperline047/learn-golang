//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import (
	"fmt"
)

type Operation byte

const (
	add Operation = iota
	sub
	mul
	div
)

func (op Operation) calculate(a, b int) int {
	switch op {
	case add:
		return a + b
	case sub:
		return a - b
	case mul:
		return a * b
	case div:
		return a / b
	default:
		panic("Invalid Operation")
	}
}

func (op Operation) calculate2(a, b int) int {
	return []int{(a + b), (a - b), (a * b), (a / b)}[op]
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50

	fmt.Println(add.calculate2(2, 2)) // = 4

	fmt.Println(sub.calculate2(10, 3)) // = 7

	fmt.Println(mul.calculate2(3, 3)) // = 9

	fmt.Println(div.calculate2(100, 2)) // = 50
}
