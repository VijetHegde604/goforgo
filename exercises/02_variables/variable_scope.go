package main

import "fmt"

var globalMessage = "Global"

func main() {
	var localMessage = "Local"
	if true {
		blockMessage := "Block"
		fmt.Printf("%s %s %s\n", globalMessage, localMessage, blockMessage)
	}

	fmt.Printf("%s %s\n", globalMessage, localMessage)
}
