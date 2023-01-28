//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssembly(title string, parts []Part) {
	fmt.Println()
	fmt.Println("----------------", title, "----------------")

	for i := 0; i < len(parts); i++ {
		part := parts[i]
		fmt.Println(part)
	}
}

func main() {
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"Washer", "Dryer", "Fryer"}
	printAssembly("Initial Parts", assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Packging", "Branding")
	printAssembly("Updated with new parts", assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	printAssembly("Only New parts", assemblyLine)
}
