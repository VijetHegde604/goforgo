package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for num := range numbers {
		sum += num
	}
	return sum
}

func printAll(words ...string) {
	for index, word := range words {
		fmt.Printf("%d:%s", index, word)
	}
}

func main() {
	sum()

	sum(1, 2, 3)

	sum(10, 20, 30, 40, 50)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	printAll("apple", "banana", "cherry")
}
