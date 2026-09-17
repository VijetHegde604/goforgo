package main

import "fmt"

func main() {
	age := 20
	temperature := 25
	score := 85

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	if temperature > 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's not hot outside")
	}

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	if x := 10 + 5; x > 12 {
		fmt.Println("x is greater than 12")
	}

	if age >= 16 {
		if age >= 21 {
			fmt.Println("Can drink alcohol")
		}
		fmt.Println("Can drive but not drink alcohol")
	} else {
		fmt.Println("Too young to drive")
	}
}
