package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}
	oldString := "Oldy"
	newString := "Newly"

	replace(&oldString, newString, &counter)
	fmt.Println(oldString, newString)

	phrase := []string{oldString, newString}
	fmt.Println(phrase)

	replace(&phrase[0], "Newer", &counter)
	fmt.Println(phrase)
}
