// slice_sorting_custom.go
// Learn custom sorting with slices using sort.Slice and sort.SliceStable
// Practice sorting complex data structures and custom comparison functions

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Age    int
	salary float64
}

func main() {
	// Create a slice of Person structs
	people := []Person{
		{Name: "Vijet", Age: 21, salary: 10000.0},
		{Name: "Vinay", Age: 22, salary: 20000.0},
		{Name: "Joe", Age: 23, salary: 30000.0},
		{Name: "Dirt", Age: 24, salary: 40000.0},
		{Name: "Dave", Age: 25, salary: 50000.0},
	}

	fmt.Println("Original people:")
	printPeople(people)

	// Sort by age (ascending) using sort.Slice
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })

	fmt.Println("\nSorted by age (ascending):")
	printPeople(people)

	// Sort by salary (descending) using sort.Slice

	fmt.Println("\nSorted by salary (descending):")
	printPeople(people)
	sort.Slice(people, func(i, j int) bool { return people[i].salary > people[j].salary })

	// Sort by name (alphabetical) using sort.Slice
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("\nSorted by name (alphabetical):")
	printPeople(people)

	// Multi-level sorting: first by age, then by salary if ages are equal
	// Use sort.SliceStable for stable sorting
	sort.SliceStable(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].salary < people[j].salary
		}
		return people[i].Age < people[j].Age
	})
	fmt.Println("\nSorted by age, then salary (stable sort):")
	printPeople(people)

	// Custom sorting with complex conditions
	// Sort by: under 30s first, then by salary descending
	sort.SliceStable(people, func(i, j int) bool {
		iUnder30 := people[i].Age < 30
		jUnder30 := people[j].Age < 30

		if iUnder30 != jUnder30 {
			return iUnder30
		}

		return people[i].salary > people[j].salary
	})

	fmt.Println("\nCustom sort (under 30s first, then by salary desc):")
	printPeople(people)

	// Sort a slice of strings by length, then alphabetically
	words := []string{"apple", "pie", "banana", "cat", "elephant", "dog", "a"}
	fmt.Printf("\nOriginal words: %v\n", words)

	// Sort by length first, then alphabetically for same length
	sort.SliceStable(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) < len(words[j])
		}
		return words[i] < words[j]
	})

	fmt.Printf("Sorted by length, then alphabetically: %v\n", words)

	// Sort a slice of integers by absolute value
	numbers := []int{-5, 3, -1, 8, -10, 2, -3}
	fmt.Printf("\nOriginal numbers: %v\n", numbers)

	// Sort by absolute value
	sort.SliceStable(numbers, func(i, j int) bool {
		return abs(numbers[i]) < abs(numbers[j])
	})
	fmt.Printf("Sorted by absolute value: %v\n", numbers)

	// Check if a slice is sorted with custom comparison
	ages := []int{25, 30, 35, 40, 45}
	isSorted := sort.SliceIsSorted(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	}) // Use sort.SliceIsSorted with custom function
	fmt.Printf("\nAges %v is sorted: %t\n", ages, isSorted)
}

// Helper function to print people slice
func printPeople(people []Person) {
	for _, p := range people {
		fmt.Printf("  %s (Age: %d, Salary: $%d)\n", p.Name, p.Age, p.salary)
	}
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
