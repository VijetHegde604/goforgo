// multidimensional_arrays.go
// Learn how to work with 2D and higher dimensional arrays

package main

import "fmt"

func main() {
	//  Declare a 3x3 matrix (2D array) of integers
	// Hint: var name [rows][cols]type
	var matrix [3][3]int // Complete this declaration

	//  Initialize a 2x4 array with values
	// Hint: arrayName := [rows][cols]type{{row1}, {row2}}
	grades := [2][4]int{{1, 2, 3, 4}, {4, 3, 2, 1}} // Complete this initialization

	//  Access and modify elements in the matrix
	// Set matrix[1][1] to 5
	// Complete this assignment
	matrix[1][1] = 5

	//  Print the matrix using nested loops
	fmt.Println("Matrix:")
	// Write nested loops to print the matrix
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(matrix[i][j], " ")
		}
	}

	//  Print the grades array using nested range loops
	fmt.Println("\nGrades:")
	// Write nested range loops
	for i := range grades {
		for j := range grades[i] {
			fmt.Print(grades[i][j], " ")
		}
		fmt.Println()
	}

	//  Calculate the sum of all elements in grades
	var totalSum int
	// Write nested loops to calculate sum
	for i := range grades {
		for j := range grades[i] {
			totalSum += grades[i][j]
		}
	}
	fmt.Println("Total sum of grades:", totalSum)

	//  Find the dimensions of the grades array
	rows := len(grades)    // Get number of rows
	cols := len(grades[0]) // Get number of columns (if rows > 0)
	fmt.Printf("Grades array dimensions: %dx%d\n", rows, cols)
}
