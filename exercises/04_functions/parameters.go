package main

import "fmt"

func greetPerson(name string) {
	fmt.Printf("Hello, %s!", name)
}

func add(a, b int) {
	fmt.Println("The sum is: ", a+b)
}
func main() {
	greetPerson("Alice")

	greetPerson("Bob")

	add(5, 3)

	add(10, 7)
}
