//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import (
	"fmt"
	"math/rand"
)

const (
	Active   = true
	Inactive = false
)

type SecureItem struct {
	name        string
	securityTag bool
}

func changeTag(item *SecureItem, state bool) {
	item.securityTag = state
}

func checkout(items []SecureItem) {
	for index := range items {
		changeTag(&items[index], Inactive)
	}
}

func main() {
	rand.Seed(4)

	shirt := SecureItem{"Soap", Active}
	hanger := SecureItem{"Hanger", Active}
	banana := SecureItem{"Banana", Active}
	detergent := SecureItem{"Detergent", Active}

	items := []SecureItem{shirt, hanger, banana, detergent}
	fmt.Println("Initial items", items)

	randomIndex := rand.Intn(4)
	changeTag(&items[randomIndex], Inactive)
	fmt.Println()
	fmt.Println("Tag", items[randomIndex].name, "deactivated")
	fmt.Println("Modified items", items)

	checkout(items)
	fmt.Println()
	fmt.Println("All tags deactivated!", items)
}
