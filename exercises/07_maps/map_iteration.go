// map_iteration.go
// Learn different ways to iterate over maps in Go

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sample data
	grades := map[string]int{
		"Alice":   92,
		"Bob":     85,
		"Charlie": 78,
		"Diana":   96,
		"Eve":     88,
	}

	inventory := map[string]int{
		"apples":  50,
		"bananas": 30,
		"oranges": 25,
	}

	// Iterate over map with both keys and values
	// Hint: for key, value := range map { ... }
	fmt.Println("Student grades:")
	// Write your iteration loop here
	for key, val := range grades {
		fmt.Printf("Student: %s, Grade: %d\n", key, val)
	}
	// Iterate over map keys only
	// Hint: for key := range map { ... }
	fmt.Println("\nStudent names (keys only):")
	// Write your iteration loop here
	for key := range grades {
		fmt.Printf("Student: %s\n", key)
	}

	// Iterate over map values only
	// Hint: for _, value := range map { ... }
	fmt.Println("\nGrade values only:")
	// Write your iteration loop here
	for _, val := range grades {
		fmt.Printf("Grades: %d\n", val)
	}
	// Calculate the sum and average of grades
	var sum int
	var count int
	// Write loop to calculate sum and count
	for _, grad := range grades {
		sum += grad
		count += 1
	}

	average := float64(sum) / float64(count)
	fmt.Printf("\nTotal sum: %d, Average: %.2f\n", sum, average)

	// Find the highest grade and student
	var maxGrade int
	var topStudent string
	// Write loop to find maximum grade and corresponding student
	maxGrade = 0
	for key, val := range grades {
		if val > maxGrade {
			maxGrade = val
			topStudent = key
		}
	}

	fmt.Printf("Top student: %s with grade %d\n", topStudent, maxGrade)

	// Iterate in sorted order by keys
	// Note: Maps are unordered, so we need to sort keys separately
	var keys []string
	// Extract keys into a slice
	for key := range inventory {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	fmt.Println("\nInventory (sorted by product name):")
	// Iterate using sorted keys
	for _, val := range keys {
		fmt.Println(inventory[val])
	}
	// Count items with specific criteria
	lowStock := 0
	threshold := 30
	// Count items with stock <= threshold
	for _, val := range inventory {
		if val <= threshold {
			lowStock += 1
		}
	}

	fmt.Printf("\nProducts with low stock (<= %d): %d\n", threshold, lowStock)
}
