package main

import "fmt"

func main() {
	fmt.Print("Start")
	goto skip
	fmt.Print("This should be skipped")
skip:
	fmt.Print("End")

	counter := 0
loop:
	fmt.Println(counter)
	counter++
	if counter < 3 {
		goto loop
	}
	fmt.Println("Done with goto loop")

	success := false
	attempt := 1
retry:
	fmt.Printf("Attempt: %d", attempt)
	if attempt < 3 && !success {
		if attempt == 2 {
			success = true
		}
		attempt++
		goto retry
	}
	if success {
		fmt.Println("Operation succeeded!")
	} else {
		fmt.Println("Operation failed after 3 attempts")
	}

	file := "data.txt"
	processed := false

	fmt.Printf("Opening file: %s", file)
	if file != "data.txt" {
		goto cleanup
	}
	fmt.Println("Processing file...")
	processed = true
cleanup:
	fmt.Println("Cleaning up...")

	if processed {
		fmt.Println("File processed successfully")
	} else {
		fmt.Println("File processing was skipped")
	}
}
