package main

import "fmt"

func main() {
	day := 3
	grade := 'B'
	score := 85

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

	switch grade {
	case 'a', 'A':
		fmt.Println("Excellence")
	case 'b', 'B':
		fmt.Println("Good")
	case 'c', 'C':
		fmt.Println("Average")
	case 'd', 'D':
		fmt.Println("Needs improvemnet")
	default:
		fmt.Println("Invalid grade")
	}

	switch {
	case score >= 90:
		fmt.Println("Outstanding")
	case score >= 80:
		fmt.Println("Very Good")
	case score >= 70:
		fmt.Println("Good")
	case score >= 60:
		fmt.Println("Satisfactory")
	default:
		fmt.Println("Needs Work")
	}

	switch day {
	case 1:
		fmt.Println("Start of work week")
		fallthrough
	case 2, 3, 4:
		fmt.Println("Weekday")
		fallthrough
	case 5:
		fmt.Println("TGIF!")
		fallthrough
	case 6, 7:
		fmt.Println("Weekend")
	}

	switch x := score / 10; x {
	case 10, 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
