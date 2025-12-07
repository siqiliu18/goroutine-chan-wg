package main

import (
	"fmt"
	"sync"
)

// Problem 4E

func main() {
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for curr := range ch {
			fmt.Println("Received: ", curr)
		}
	}()

	for i := 1; i <= 5; i++ {
		ch <- i
		if len(ch) == cap(ch) {
			fmt.Printf("Sending %v (blocks until receiver takes one)\n", i)
		} else {
			fmt.Printf("Sending %v (non-blocking, buffer has space)\n", i)
		}
	}

	close(ch)
	wg.Wait()
}
