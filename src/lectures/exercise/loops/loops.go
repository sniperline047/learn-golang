//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func divisibleByThree(number int) bool {
	return number%3 == 0
}

func divisibleByFive(number int) bool {
	return number%5 == 0
}

func main() {
	for i := 1; i < 50; i++ {
		if divisibleByThree(i) && divisibleByFive(i) {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FizzBuzz")
		} else if divisibleByThree(i) {
			//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println("Fizz")
		} else if divisibleByFive(i) {
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
