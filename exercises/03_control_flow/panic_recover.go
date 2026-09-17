package main

import "fmt"

func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func processArray(array []int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from array access panic")
		}
	}()
	val := array[10]
	fmt.Println(val)
}

func demonstratePanic() {
	fmt.Println("Before panic")
	panic("Something went wrong")
	fmt.Println("After panic")
}

func nestedPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Outer recovery:", r)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Inner recovery:", r)
			}
		}()
		panic("Inner panic")
	}()
}

func main() {

	safeDivide(10, 2)

	safeDivide(10, 0)

	processArray([]int{1, 2, 3})

	processArray([]int{1, 2, 3, 4, 5, 6, 6, 7, 5, 6, 7, 7, 8, 9, 8})

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic")
			}
		}()
		demonstratePanic()
	}()

	nestedPanic()

	fmt.Println("Program continues normally")
}
