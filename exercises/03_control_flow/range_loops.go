package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fruits := []string{"apple", "banana", "cherry"}
	message := "Hello"
	scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}

	for i, val := range numbers {
		fmt.Printf("Index %d: %d", i, val)
	}

	for _, val := range fruits {
		fmt.Printf("Fruit: %s", val)
	}

	for i, _ := range fruits {
		fmt.Printf("Index: %d", i)
	}

	for i, char := range message {
		fmt.Printf("Position %d: %c", i, char)
	}

	for key, val := range scores {
		fmt.Printf("%s scored %d", key, val)
	}

	for key, _ := range scores {
		fmt.Printf("Student: %s", key)
	}

	for _, val := range scores {
		fmt.Printf("Score: %s", val)
	}

	ch := make(chan int, 3)
	for i := range 3 {
		ch <- i
	}
	close(ch)

	for j := range ch {
		fmt.Printf("Received: %d", j)
	}
}
