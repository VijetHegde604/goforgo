package main

import "fmt"

// Define a function called 'makeCounter' that returns a function
// The returned function should have no parameters and return an int
// Use a closure to create a counter that increments each time it's called
// The counter should start at 0 and increment by 1 each call
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Define a function called 'makeMultiplier' that takes an int parameter 'factor'
// It should return a function that takes an int and returns an int
// The returned function should multiply its input by the factor
func makeMultiplier(factor int) func(int) int {
	return func(input int) int {
		return input * factor
	}
}

func main() {
	//  Create a counter using makeCounter
	counter1 := makeCounter()

	//  Call the counter 3 times and print each result
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	//  Create another counter using makeCounter
	counter2 := makeCounter()

	//  Call the second counter 2 times and print each result
	// Note: each counter should maintain its own state
	fmt.Println(counter2())
	fmt.Println(counter2())

	//  Create a multiplier that multiplies by 3 using makeMultiplier
	multiplier := makeMultiplier(3)

	//  Use the multiplier on 5 and on 10, print the results
	fmt.Println(multiplier(5))
	fmt.Println(multiplier(10))

	//  Create a multiplier that multiplies by 7 using makeMultiplier
	multiplier2 := makeMultiplier(7)

	//  Use the second multiplier on 4 and on 6, print the results
	fmt.Println(multiplier2(4))
	fmt.Println(multiplier2(6))
}
