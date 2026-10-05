package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	wg := sync.WaitGroup{}

	counter := SafeCounter{}

	for range 100 {
		wg.Go(func() {
			counter.Increment()
		})
	}
	wg.Wait()
	fmt.Println("Count: ", counter.Value())
}
