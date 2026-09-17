package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for num := range numbers {
		if num > 5 {
			fmt.Printf("Found number greater than 5: %d", num)
			break
		}
	}

	for num := range numbers {
		if num%2 != 0 {
			continue
		}
		fmt.Print(num)
	}

OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print(i + j)

			if (i + j) >= 4 {
				break OuterLoop
			}
		}
	}

NextI:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == j {
				continue NextI
			}
			fmt.Printf("i=%d, j=%d", i, j)
		}
	}

Candidates:
	for i := 2; i <= 20; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue Candidates
			}
		}
		fmt.Println(i)
	}

	secret := 7
	guesses := []int{3, 7, 5, 9, 7, 2}
	for guess := range guesses {
		if guess == secret {
			break
		}
		fmt.Print(guess)
		continue
	}
}
