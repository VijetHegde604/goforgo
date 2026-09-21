// slice_tricks.go
// Learn advanced slice operations and common patterns

package main

import (
	"fmt"
	"slices"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//  Remove element at index 3 (value 4)
	// Use append to combine slices before and after the index
	index := 3
	// Hint: append(slice[:index], slice[index+1:]...)
	data = append(data[:index], data[index+1:]...) // Complete this removal operation

	fmt.Println("After removing index 3:", data)

	//  Insert element 99 at index 2
	// Split the slice and insert the new element
	insertIndex := 2
	insertValue := 99
	// Hint: append(append(slice[:index], value), slice[index:]...)
	data = append(append(data[:insertIndex], insertValue), data[insertIndex+1:]...) // Complete this insertion operation

	fmt.Println("After inserting 99 at index 2:", data)

	//  Reverse a slice in place
	numbers := []int{1, 2, 3, 4, 5}
	// Use two pointers approach
	// Complete the reversal logic
	for i := range len(numbers) / 2 {
		numbers[i], numbers[len(numbers)-1-i] = numbers[len(numbers)-1-i], numbers[i]
	}
	fmt.Println("Reversed numbers:", numbers)

	//  Find and remove all occurrences of a value
	values := []int{1, 2, 3, 2, 4, 2, 5}
	target := 2

	// Filter out the target value
	var filtered []int
	// Write a loop to filter out target value
	for _, val := range values {
		if val == target {
			continue
		}
		filtered = append(filtered, val)
	}

	fmt.Printf("After removing all %d: %v\n", target, filtered)

	//  Check if slice contains a value
	haystack := []string{"apple", "banana", "cherry", "date"}
	needle := "cherry"

	var found bool
	// Write logic to check if needle exists in haystack
	found = slices.Contains(haystack, needle)

	if found {
		fmt.Printf("Found '%s' in slice\n", needle)
	} else {
		fmt.Printf("'%s' not found in slice\n", needle)
	}

	//  Get unique elements from a slice
	duplicates := []int{1, 2, 2, 3, 1, 4, 3, 5}
	var unique []int

	// Use a map to track seen elements
	// Write logic to get unique elements
	seen := make(map[int]bool)
	for _, val := range duplicates {
		if !seen[val] {
			seen[val] = true
			unique = append(unique, val)
		}
	}

	fmt.Println("Unique elements:", unique)
}
