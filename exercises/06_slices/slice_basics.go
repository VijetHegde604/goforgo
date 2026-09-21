// slice_basics.go
// Learn the fundamentals of slices in Go
// Slices are dynamic arrays with flexible size

package main

import "fmt"

func main() {
	//  Create a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // Create a slice from arr containing elements 1-3

	//  Create a slice using make()
	// Hint: make([]type, length, capacity)
	slice2 := make([]string, 3, 5) // Create a slice of 3 strings with capacity 5

	//  Create a slice using slice literal
	// Hint: []type{values}
	slice3 := []int{10, 20, 30} // Create a slice with numbers 10, 20, 30

	//  Create an empty slice
	var slice4 []int // Complete this declaration

	//  Check if slice4 is nil
	if len(slice4) == 0 {
		fmt.Println("slice4 is nil")
	}

	//  Get length and capacity of slice2
	length := len(slice2)   // Get length of slice2
	capacity := cap(slice2) // Get capacity of slice2

	//  Modify elements in slice1
	// Change the second element to 99
	// Complete this assignment
	slice1[1] = 99

	// Print results
	fmt.Println("Original array:", arr)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Printf("slice2 - length: %d, capacity: %d\n", length, capacity)
}
