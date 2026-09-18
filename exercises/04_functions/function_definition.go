package main

import "fmt"

func greet() {
	fmt.Println("Hello from a function!")
}

func main() {
	fmt.Println("This runs before the function call")
	fmt.Println("This runs after the function call")
	greet()
}
