package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	items []string
}

func (s *Stuff) Get(index int) (string, error) {
	if index >= len(s.items) {
		return "", errors.New(fmt.Sprintf("Invalid index %v given, there are only %v items", index, len(s.items)))
	} else {
		return s.items[index], nil
	}
}

func main() {
	items := Stuff{[]string{"Keys", "Phone", "Laptop", "Watch"}}

	item, err := items.Get(69)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Item:", item)
	}
}
