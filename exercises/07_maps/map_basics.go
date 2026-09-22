// map_basics.go
// Learn the fundamentals of maps in Go
// Maps are key-value pairs similar to hash tables or dictionaries

package main

import "fmt"

func main() {
	// Declare a map using make()
	// Hint: make(map[keyType]valueType)
	ages := make(map[string]int) // Create a map from string to int

	// Initialize a map with values using map literal
	// Hint: map[keyType]valueType{key: value, key: value}
	colors := map[string]string{"red": "#FF0000", "blue": "#275BF5", "green": "#3FF527"} // Create a map with color names as keys and hex codes as values

	// Add elements to the ages map
	ages["Alice"] = 25
	ages["Bob"] = 30
	// Add "Charlie" with age 35

	// Access a value from the map
	aliceAge := ages["Alice"] // Get Alice's age from the map

	// Check if a key exists using the comma ok idiom
	// Hint: value, ok := map[key]
	bobAge, exists := ages["Bob"] // Check if Bob exists in ages map

	if exists {
		fmt.Printf("Bob's age: %d\n", bobAge)
	}

	// Try to access a non-existent key
	// Hint: Non-existent keys return zero value
	unknown := ages["Unknown"] // Try to get "Unknown" from ages map
	fmt.Printf("Unknown person's age (zero value): %d\n", unknown)

	// Delete a key from the map
	// Hint: delete(map, key)
	// Delete "Alice" from ages map
	delete(ages, "Alice")

	// Get the length of the map
	length := len(ages) // Get length of ages map

	// Create a nil map and try to assign (this will panic if uncommented)
	var nilMap map[string]int
	// nilMap["test"] = 1 // This would panic!

	// Check if a map is nil
	if nilMap == nil {
		fmt.Println("nilMap is nil - cannot assign to it")
	}

	// Print results
	fmt.Println("Ages map:", ages)
	fmt.Println("Colors map:", colors)
	fmt.Printf("Alice's age: %d\n", aliceAge)
	fmt.Printf("Map length: %d\n", length)
}
