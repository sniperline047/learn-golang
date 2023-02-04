package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	//            :key:    :value:

	// insertion
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 10
	shoppingList["bread"] += 1

	// updation
	shoppingList["egss"] += 1
	fmt.Println(shoppingList)

	// deletetion
	delete(shoppingList, "random")
	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, New shopping list", shoppingList)

	// selection
	fmt.Println("need", shoppingList["eggs"], "eggs")

	// check for key using 2nd arg
	cereal, found := shoppingList["cereal"]

	if !found {
		fmt.Println("Cereal not found!")
		shoppingList["cereal"] = 10
		fmt.Println("Adding", shoppingList["cereal"], "cereal to shopping list")
	} else {
		fmt.Println(cereal, "Cereal are in the shopping list")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "in the shopping list")
}
