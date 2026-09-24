// interface_assertion.go
// Learn advanced type assertion patterns and interface checking

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Stringer interface {
	String() string
}

type Counter interface {
	Count() int
}

type Resetter interface {
	Reset()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s and Age: %d", p.Name, p.Age)
}

// WordCounter implements Stringer, Counter, and Resetter.
type WordCounter struct {
	words int
}

func (wc *WordCounter) AddWords(text string) {
	if text == "" {
		return
	}

	words := strings.Fields(text)
	wc.words += len(words)
}

func (wc WordCounter) String() string {
	return fmt.Sprintf("Word Count: %d", wc.words)
}

func (wc WordCounter) Count() int {
	return wc.words
}

func (wc *WordCounter) Reset() {
	wc.words = 0
}

// checkStringer checks whether a value implements Stringer.
func checkStringer(value interface{}) {
	if stringer, ok := value.(Stringer); ok {
		fmt.Printf("✓ Implements Stringer: %s\n", stringer.String())
	} else {
		fmt.Printf("✗ Does not implement Stringer: %v\n", value)
	}
}

// checkInterfaces checks whether a value implements multiple interfaces.
func checkInterfaces(value interface{}) {
	fmt.Printf("Checking interfaces for %T:\n", value)

	if _, ok := value.(Stringer); ok {
		fmt.Println("  ✓ Implements Stringer")
	} else {
		fmt.Println("  ✗ Does not implement Stringer")
	}

	if _, ok := value.(Counter); ok {
		fmt.Println("  ✓ Implements Counter")
	} else {
		fmt.Println("  ✗ Does not implement Counter")
	}

	if _, ok := value.(Resetter); ok {
		fmt.Println("  ✓ Implements Resetter")
	} else {
		fmt.Println("  ✗ Does not implement Resetter")
	}
}

// convertToString converts different types to strings.
func convertToString(value interface{}) string {
	// Prefer the Stringer interface.
	if stringer, ok := value.(Stringer); ok {
		return stringer.String()
	}

	// Handle basic types.
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// processData conditionally calls methods based on implemented interfaces.
func processData(data interface{}) {
	fmt.Printf("Processing %T: ", data)
	fmt.Printf("Value = %v", data)

	if stringer, ok := data.(Stringer); ok {
		fmt.Printf(", String() = %s", stringer.String())
	}

	if counter, ok := data.(Counter); ok {
		fmt.Printf(", Count() = %d", counter.Count())
	}

	fmt.Println()
}

func main() {
	// Create test objects.
	person := Person{
		Name: "Alice",
		Age:  30,
	}

	wc := WordCounter{}
	wc.AddWords("hello world from Go")

	plainInt := 42
	plainString := "just a string"

	// --------------------------------------------------
	// Stringer Interface Check
	// --------------------------------------------------

	fmt.Println("=== Stringer Interface Check ===")

	checkStringer(person)
	checkStringer(&wc)
	checkStringer(plainInt)
	checkStringer(plainString)

	// --------------------------------------------------
	// Multiple Interface Check
	// --------------------------------------------------

	fmt.Println("\n=== Multiple Interface Check ===")

	checkInterfaces(person)
	checkInterfaces(&wc)
	checkInterfaces(plainInt)

	// --------------------------------------------------
	// String Conversion
	// --------------------------------------------------

	fmt.Println("\n=== String Conversion ===")

	values := []interface{}{
		person,
		&wc,
		123,
		3.14159,
		true,
		[]int{1, 2, 3},
	}

	for _, value := range values {
		str := convertToString(value)
		fmt.Printf("%T -> %q\n", value, str)
	}

	// --------------------------------------------------
	// Conditional Method Calls
	// --------------------------------------------------

	fmt.Println("\n=== Conditional Method Calls ===")

	processData(person)
	processData(&wc)
	processData(plainInt)
	processData(plainString)

	// --------------------------------------------------
	// Interface Nil Check
	// --------------------------------------------------

	fmt.Println("\n=== Interface Nil Check ===")

	var nilStringer Stringer
	var nilPerson *Person

	if nilStringer == nil {
		fmt.Println("nilStringer is nil")
	}

	// An interface containing a nil pointer is NOT itself nil.
	nilStringer = nilPerson

	if nilStringer == nil {
		fmt.Println("nilStringer with nil pointer is nil")
	} else {
		fmt.Printf(
			"nilStringer with nil pointer is not nil: %T\n",
			nilStringer,
		)

		// Safely detect the typed-nil pointer.
		value := reflect.ValueOf(nilStringer)

		if value.Kind() == reflect.Ptr && value.IsNil() {
			fmt.Println("Underlying pointer is nil, so String() will not be called")
		} else {
			fmt.Println(nilStringer.String())
		}
	}

	// --------------------------------------------------
	// Counter Operations
	// --------------------------------------------------

	fmt.Println("\n=== Counter Operations ===")

	// Type assertions work on interface values.
	var value interface{} = wc

	if counter, ok := value.(Counter); ok {
		fmt.Printf("Initial count: %d\n", counter.Count())

		wc.AddWords("more words here")
		fmt.Printf("After adding words: %d\n", counter.Count())

		if resetter, ok := value.(Resetter); ok {
			resetter.Reset()
			fmt.Printf("After reset: %d\n", counter.Count())
		}
	}
}
