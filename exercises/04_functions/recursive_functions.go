package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func countdown(n int) {
	if n == 0 {
		fmt.Println("Blast off!")
		return
	}
	fmt.Println(n)
	countdown(n - 1)
}

func main() {

	fmt.Println(factorial(5))

	fmt.Println(factorial(0))

	fmt.Println(fibonacci(7))

	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))

	countdown(5)
}
