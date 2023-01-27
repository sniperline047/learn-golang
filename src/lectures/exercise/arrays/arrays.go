//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
func printStats(list [4]Product) {
	fmt.Println("Printing product list stats")

	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]

		if item.price > 0 && item.name != "" {
			cost += item.price
			totalItems++
		}
	}

	fmt.Println(list[totalItems-1].name, "is the last item")
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", cost)
}

func main() {
	products := [4]Product{
		{name: "Shampoo", price: 69},
		{name: "Body Wash", price: 420},
		{name: "Deodrant", price: 911},
	}

	fmt.Println("Current product list", products)
	printStats(products)

	fmt.Println("Adding product to list")
	products[3] = Product{name: "Pen", price: 112}
	printStats(products)
}
