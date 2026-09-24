package main

import (
	"fmt"
	"math"
	"strings"
)

// Function that accepts any type using an empty interface.
func printAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// Function that handles different types using type assertions.
func handleValue(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Printf("It's a string: %s (length: %d)\n", str, len(str))
		return
	}

	if num, ok := value.(int); ok {
		fmt.Printf("It's an int: %d (doubled: %d)\n", num, num*2)
		return
	}

	if b, ok := value.(bool); ok {
		fmt.Printf("It's a bool: %t (negated: %t)\n", b, !b)
		return
	}

	fmt.Printf("Unknown type: %T\n", value)
}

// Function using a type switch.
func processValue(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf(
			"String processing: '%s' -> '%s'\n",
			v,
			strings.ToUpper(v),
		)

	case int:
		fmt.Printf(
			"Int processing: %d -> %d\n",
			v,
			v*v,
		)

	case float64:
		fmt.Printf(
			"Float processing: %.2f -> %.2f\n",
			v,
			math.Sqrt(v),
		)

	case []int:
		var sum int

		for _, num := range v {
			sum += num
		}

		fmt.Printf(
			"Slice processing: %v -> sum: %d\n",
			v,
			sum,
		)

	case bool:
		if v {
			fmt.Printf("Bool processing: %t -> yes\n", v)
		} else {
			fmt.Printf("Bool processing: %t -> no\n", v)
		}

	default:
		fmt.Printf("Unknown type processing: %T = %v\n", v, v)
	}
}

// Function that stores mixed types in a slice.
func demonstrateSliceOfInterface() {
	mixed := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
	}

	fmt.Println("Mixed slice contents:")

	for i, item := range mixed {
		fmt.Printf("Index %d: %v (%T)\n", i, item, item)
	}
}

// Function that uses interface{} in a map.
func demonstrateMapOfInterface() {
	data := make(map[string]interface{})

	data["name"] = "Alice"
	data["age"] = 30
	data["height"] = 5.8
	data["married"] = true

	fmt.Println("Map with mixed value types:")

	for key, value := range data {
		fmt.Printf("%s: %v (%T)\n", key, value, value)
	}

	// Safely extract values using type assertions.
	if name, ok := data["name"].(string); ok {
		fmt.Printf("Name is: %s\n", name)
	}

	if age, ok := data["age"].(int); ok {
		fmt.Printf("Age is: %d\n", age)
	}
}

func main() {
	fmt.Println("=== Empty Interface Demo ===")

	printAnything("hello world")

	fmt.Println("\n=== Type Assertion Demo ===")

	handleValue(3.14)

	fmt.Println("\n=== Type Switch Demo ===")

	processValue([]int{1, 23, 4, 5})

	fmt.Println("\n=== Slice of Interface Demo ===")

	demonstrateSliceOfInterface()

	fmt.Println("\n=== Map of Interface Demo ===")

	demonstrateMapOfInterface()

	fmt.Println("\n=== Safe vs Unsafe Type Assertion ===")

	var value interface{} = "hello"

	// Safe type assertion.
	if str, ok := value.(string); ok {
		fmt.Printf("Safe assertion: %s\n", str)
	}

	// Unsafe type assertion.
	// This works because value actually contains a string.
	str := value.(string)
	fmt.Printf("Unsafe assertion: %s\n", str)

	// This would panic because value contains a string, not an int.
	// num := value.(int)
}
