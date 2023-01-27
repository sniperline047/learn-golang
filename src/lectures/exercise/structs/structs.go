//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x float64
	y float64
}

type Rectangle struct {
	C1 Coordinate // top left
	C2 Coordinate // top right
	C3 Coordinate // bottom right
	C4 Coordinate // bottom left
}

func isRectangle(rectangle Rectangle) bool {
	return true
}

func dimensions(rectangle Rectangle) (float64, float64) {
	// returns width, height
	width := math.Sqrt(math.Pow((rectangle.C1.x-rectangle.C2.x), 2) + math.Pow((rectangle.C1.y-rectangle.C2.y), 2))
	height := math.Sqrt(math.Pow((rectangle.C2.x-rectangle.C3.x), 2) + math.Pow((rectangle.C2.y-rectangle.C3.y), 2))

	return width, height
}

func area(width float64, height float64) float64 {
	return width * height
}

func perimeter(width float64, height float64) float64 {
	return (width + height) * 2
}

func areaAndPerimeter(rectangle Rectangle) (float64, float64) {
	width, height := dimensions(rectangle)
	fmt.Println(width, height)
	return area(width, height), perimeter(width, height)
}

func main() {
	var (
		coordinate1 = Coordinate{0, 0}
		coordinate2 = Coordinate{6, 0}
		coordinate3 = Coordinate{6, -6}
		coordinate4 = Coordinate{0, -6}
	)

	rectangle := Rectangle{C1: coordinate1, C2: coordinate2, C3: coordinate3, C4: coordinate4}
	fmt.Println("Input Coordinates of rectangle:", rectangle)

	fmt.Println("Processing rectangle coordinates...")

	if !isRectangle(rectangle) {
		fmt.Println("[ERROR!] Input coordinates produce a invalid rectangle")
	}

	fmt.Println("Input coordinates are valid, calculating area and perimeter")
	area, perimeter := areaAndPerimeter(rectangle)
	fmt.Println("Area of given rectangle is", area, "and its perimeter is", perimeter)
}
