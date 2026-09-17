package main

import "fmt"

func describeType(input any) {
	switch in := input.(type) {
	case string:
		fmt.Println("It's a string with length:", len(in))
	case int:
		fmt.Println("It's an integer with value:", in)
	case bool:
		fmt.Println("It's a boolean with value:", in)
	case []int:
		fmt.Println("It's an int slice with length:", len(in))
	default:
		fmt.Println("Unknown type")
	}
}

func processValue(val interface{}) {
	switch v := val.(type) {
	case string:
		if len(v) > 5 {
			fmt.Println("Long string:", v)
		} else {
			fmt.Println("Short string:", v)
		}
	case int:
		if v > 0 {
			fmt.Println("Positive:", v)
		} else {
			fmt.Println("Non-positive:", v)
		}
	case float64:
		fmt.Printf("Float with 2 decimal places: %.2f", v)
	default:
		fmt.Printf("Cannot process type: %T", val)
	}
}

func handleError(val any) {
	switch err := val.(type) {
	case string:
		fmt.Println("String error:", err)
	case error:
		fmt.Println("Error type:", err.Error())
	case nil:
		fmt.Println("No error")
	default:
		fmt.Println("Unknown error type")
	}
}

func main() {
	describeType("hello world")

	processValue([]string{"not", "handled"})

	handleError(nil)

	func(val any) {
		switch v := val.(type) {
		case int, int32, int64:
			fmt.Println("Its int type", v)
		}
	}(10)

}
