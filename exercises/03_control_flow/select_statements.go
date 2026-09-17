package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Recieved from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Recieved from ch2:", msg2)
	}

	ch3 := make(chan int)
	select {
	case msg3 := <-ch3:
		fmt.Println("Recieved from ch3:", msg3)
	default:
		fmt.Println("No data available")
	}

	slow := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		slow <- "Message from slow"
	}()
	select {
	case msg4 := <-slow:
		fmt.Println("Recieved from slow:", msg4)
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out")
	}

	for range 5 {
		fast1 := make(chan int, 1)
		fast2 := make(chan int, 1)
		fast1 <- 1
		fast2 <- 2
		select {
		case msgf1 := <-fast1:
			fmt.Println("Recieved from fast1:", msgf1)
		case msgf2 := <-fast2:
			fmt.Println("Recieved from fast2:", msgf2)
		}
	}

	send1 := make(chan string, 1)
	send2 := make(chan string, 1)
	select {
	case send1 <- "Message1":
		fmt.Println("Sent to send1: Message1")
	case send2 <- "Message2":
		fmt.Println("Sent to send2: Message2")
	}
	select {
	case msg := <-send1:
		fmt.Println("Recieved from send1:", msg)
	case msg := <-send2:
		fmt.Println("Recieved from send2:", msg)
	}
}
