package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

// Problem 4A
func producer(id int, shared chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		val := rand.IntN(100)
		fmt.Printf("Producer %v: sent %v\n", id, val)
		shared <- val
	}
}

func main() {
	shared := make(chan int)
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Step 1: Start consumer FIRST (ready to receive)
	consumerWg.Add(1)
	go func() {
		defer consumerWg.Done()
		for val := range shared {
			fmt.Printf("Consumer: received %v\n", val)
		}
	}()

	// Step 2: Start producers (consumer is ready!)
	producerWg.Add(2)
	go producer(1, shared, &producerWg)
	go producer(2, shared, &producerWg)

	// Step 3: Wait for all producers to finish
	producerWg.Wait()

	// Step 4: Close channel (signals consumer to stop)
	close(shared)

	// Step 5: Wait for consumer to finish
	consumerWg.Wait()

	fmt.Println("All done!")
}
