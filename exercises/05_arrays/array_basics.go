// array_basics.go
// Learn the fundamentals of arrays in Go
// Arrays have a fixed size and all elements must be of the same type

package main

import "fmt"

func main() {
	// : Declare an array of 5 integers
	// Hint: var arrayName [size]type
	var numbers [5]int // Complete this declaration

	// : Initialize an array with values using array literal
	// Hint: arrayName := [size]type{value1, value2, ...}
	colors := [3]string{"Red", "Green", "Blue"} // Complete this initialization

	// : Use the ... operator to let Go determine the array size
	// Hint: arrayName := [...]type{values}
	days := []string{"Sunday", "Monday", "Tuesday"} // Complete this initialization

	// : Access the third element (index 2) of the numbers array
	// and assign it the value 42
	// Complete this assignment
	numbers[2] = 42
	// : Print the length of the colors array
	// Hint: use len() function
	fmt.Println("Length of colors array:", len(colors)) // Complete this line

	// Print all arrays to see the results
	fmt.Println("Numbers:", numbers)
	fmt.Println("Colors:", colors)
	fmt.Println("Days:", days)
}
