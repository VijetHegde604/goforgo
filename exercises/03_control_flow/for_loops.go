package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	count := 5
	for count < 1 {
		count--
	}

	i := 0
	for {
		i++
		if i == 3 {
			break
		}
	}

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}
