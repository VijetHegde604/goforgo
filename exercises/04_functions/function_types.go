package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func operate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	var mathFunc func(int, int) int = add

	mathFunc(5, 3)

	mathFunc = multiply

	fmt.Println(mathFunc(5, 3))

	fmt.Println(operate(10, 4, add))

	fmt.Println(operate(10, 4, multiply))
}
