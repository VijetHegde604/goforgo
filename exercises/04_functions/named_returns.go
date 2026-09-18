package main

func calculateArea(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func swapStrings(a, b string) (first, second string) {
	first = b
	second = a
	return
}

func main() {
	calculateArea(5.0, 3.0)
	swapStrings("hello", "world")
}
