package main

import (
	"fmt"
	"sync"
)

// Problem 4F

func sender(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for curr := range ch {
		fmt.Println("Received: ", curr)
	}
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	// wgSender := &sync.WaitGroup{}
	// wgReceiver := &sync.WaitGroup{}

	wg.Add(1)
	go sender(ch, wg)

	wg.Add(1)
	go receiver(ch, wg)

	// wgReceiver.Wait()
	// wgSender.Wait()
	wg.Wait()

	fmt.Println("All done!")
}
