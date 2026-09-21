// array_iteration.go
// Learn different ways to iterate over arrays in Go

package main

import "fmt"

func main() {
	scores := [5]int{85, 92, 78, 96, 88}
	names := [...]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

	//  Use a traditional for loop to print all scores
	// Hint: for i := 0; i < len(array); i++ { ... }
	fmt.Println("Scores using traditional for loop:")
	// Write your loop here
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	//  Use range to iterate over the names array
	// Print both index and value
	// Hint: for index, value := range array { ... }
	fmt.Println("\nNames with indices using range:")
	// Write your range loop here
	for i, val := range names {
		fmt.Printf("Index: %d and name: %s\n", i, val)
	}

	//  Use range to iterate over scores array
	// but only use the values (ignore the index)
	// Hint: for _, value := range array { ... }
	fmt.Println("\nScores using range (values only):")
	// Write your range loop here
	for _, val := range scores {
		fmt.Println(val)
	}
	//  Find the maximum score using iteration
	maxScore := scores[0] // Initialize with first element
	// Write a loop to find the maximum
	for i := range scores {
		if scores[i] > maxScore {
			maxScore = scores[i]
		}
	}
	fmt.Println("\nMaximum score:", maxScore)

	//  Calculate the average score
	var sum int
	// Write a loop to calculate sum
	for i := range scores {
		sum += scores[i]
	}
	average := sum / len(scores) // Calculate average
	fmt.Printf("Average score: %.2f\n", average)
}
