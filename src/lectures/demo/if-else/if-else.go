package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 1, 2, 3
	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 is higher than quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 is higher than Quiz 1")
	} else {
		fmt.Println("Both Quiz 1 and 2 are equal")
	}

	if average(quiz1, quiz2, quiz3) > 2 {
		fmt.Println("Better than average marks")
	} else {
		fmt.Println("Below average marks")
	}
}
