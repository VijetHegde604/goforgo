package main

import (
	"fmt"
	"math"
)

// Shape defines anything that can calculate an area.
type Shape interface {
	Area() float64
}

// Rectangle implements Shape.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements Shape.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Writer defines anything that can write data.
type Writer interface {
	Write(data []byte) (int, error)
}

// FileWriter implements Writer.
type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing %d bytes to file: %s\n", len(data), fw.Filename)
	fmt.Printf("Content: %s\n", string(data))
	return len(data), nil
}

// ConsoleWriter implements Writer.
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Printf("Console: %s", string(data))
	return len(data), nil
}

// Works with any type that implements Shape.
func printShapeInfo(s Shape) {
	fmt.Printf("Shape area: %.2f\n", s.Area())
}

// Works with any type that implements Writer.
func writeMessage(w Writer, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// -------------------------
	// Shapes
	// -------------------------

	rect := Rectangle{
		Width:  5,
		Height: 3,
	}

	circle := Circle{
		Radius: 4,
	}

	fmt.Println("Using shapes through interface:")

	printShapeInfo(rect)
	printShapeInfo(circle)

	// A slice containing different types
	// because both implement Shape.
	shapes := []Shape{
		rect,
		circle,
	}

	var totalArea float64

	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	fmt.Printf("Total area: %.2f\n", totalArea)

	// -------------------------
	// Writers
	// -------------------------

	fileWriter := FileWriter{
		Filename: "output.txt",
	}

	consoleWriter := ConsoleWriter{}

	fmt.Println("\nUsing writers through interface:")

	writeMessage(fileWriter, "Hello, File!\n")
	writeMessage(consoleWriter, "Hello, Console!\n")

	// A slice containing different types
	// because both implement Writer.
	writers := []Writer{
		fileWriter,
		consoleWriter,
	}

	for i, writer := range writers {
		message := fmt.Sprintf("Message %d\n", i+1)
		writeMessage(writer, message)
	}

	// -------------------------
	// Interface assignment
	// -------------------------

	var shape Shape

	shape = rect
	fmt.Printf("\nRectangle area: %.2f\n", shape.Area())

	shape = circle
	fmt.Printf("Circle area: %.2f\n", shape.Area())
}
