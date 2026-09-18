package main

import "fmt"

func double(value int) int {
	return value * 2
}

func getGreeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0.0, false
	} else {
		return a / b, true
	}
}

func main() {
	fmt.Println(double(5))

	fmt.Println(getGreeting("World"))

	fmt.Println(divide(10.0, 2.0))

	fmt.Println(divide(10.0, 0.0))
}
