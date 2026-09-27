// error_basics.go
// Learn the fundamentals of error handling in Go
// Go uses explicit error values rather than exceptions

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

// Function that uses errors.New()
func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	if age > 150 {
		return errors.New("age cannot exceed 150")
	}

	return nil
}

// Function that returns multiple types of errors
func parseAndValidateAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, err
	}

	// Then validate the parsed age
	err = validateAge(age)
	if err != nil {
		return 0, err
	}

	return age, nil
}

// Function that demonstrates error propagation
func processUserAge(ageStr string) {
	age, err := parseAndValidateAge(ageStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Valid age: %d\n", age)
}

// Function with named return values for cleaner error handling
func safeDivide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}

	result = a / b
	return
}

// Function that checks specific error types
func handleStringConversion(s string) {
	num, err := strconv.Atoi(s)
	if err != nil {
		// Check if the error is a *strconv.NumError
		var numErr *strconv.NumError

		if errors.As(err, &numErr) {
			fmt.Printf("Number parsing error: %v\n", err)
		} else {
			fmt.Printf("Other error: %v\n", err)
		}

		return
	}

	fmt.Printf("Converted number: %d\n", num)
}

func main() {
	fmt.Println("=== Basic Error Handling ===")

	// Test divide function
	result, err := divide(10, 2)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// Test divide by zero
	result, err = safeDivide(10, 0)

	if err != nil {
		fmt.Printf("Division by zero error: %v\n", err)
	}

	fmt.Println("\n=== Error Validation ===")

	// Test age validation
	ages := []int{25, -5, 200, 0, 45}

	for _, age := range ages {
		err := validateAge(age)

		if err != nil {
			fmt.Printf("Age %d is invalid: %v\n", age, err)
		} else {
			fmt.Printf("Age %d is valid\n", age)
		}
	}

	fmt.Println("\n=== Error Propagation ===")

	// Test parseAndValidateAge with different inputs
	testInputs := []string{"25", "-5", "abc", "200", "30"}

	for _, input := range testInputs {
		fmt.Printf("Processing '%s': ", input)
		processUserAge(input)
	}

	fmt.Println("\n=== Named Return Values ===")

	// Test safeDivide
	result, err = safeDivide(15, 3)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", result)
	}

	// Test with zero
	result, err = safeDivide(10, 0)

	if err != nil {
		fmt.Printf("Safe divide error: %v\n", err)
	}

	fmt.Println("\n=== Error Type Checking ===")

	// Test string conversion with different inputs
	testStrings := []string{"123", "abc", "456", "12.34"}

	for _, s := range testStrings {
		fmt.Printf("Converting '%s': ", s)
		handleStringConversion(s)
	}

	fmt.Println("\n=== Error Patterns ===")

	// Demonstrate ignoring errors
	// (not recommended unless you intentionally don't care about the error)
	num, _ := strconv.Atoi("123")
	fmt.Printf("Ignoring error result: %d\n", num)

	// Demonstrate multiple error checks
	s1, s2 := "10", "5"

	n1, err1 := strconv.Atoi(s1)
	n2, err2 := strconv.Atoi(s2)

	if err1 != nil || err2 != nil {
		fmt.Printf("Parsing errors: err1=%v, err2=%v\n", err1, err2)
		return
	}

	// Both parsing operations succeeded
	result, err = divide(float64(n1), float64(n2))

	if err != nil {
		fmt.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("%s / %s = %.2f\n", s1, s2, result)
	}
}
