package main

import "fmt"

func cleanup() {
	fmt.Println("CLeaning up resources")
}

func openfile(filename string) {
	fmt.Printf("Opening file: %s", filename)
	defer cleanup()
	fmt.Println("File operations complete")
}

func deferOrder() {
	fmt.Println("Start")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("End")

}

func deferWithArgs() {
	x := 10
	defer fmt.Println("Deferred x:", x)
	x = 20
	fmt.Println("Current x:", x)
}

func deferInLoop() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Loop defer:", i)
		fmt.Println("Loop iteration:", i)
	}
}

func main() {

	openfile("data.txt")

	deferOrder()

	deferWithArgs()

	deferInLoop()

	message := "Hello"
	defer fmt.Println("Deferred message:", message)
	message = "Goodbye"
	fmt.Println("Current message:", message)
}
