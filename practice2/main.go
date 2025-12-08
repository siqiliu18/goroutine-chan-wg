package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Problem 3

func sender(taskCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		taskCh <- i
	}
	close(taskCh)
}

func worker(id int, taskCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := range taskCh {
		sleep := rand.IntN(401) + 100
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker %v processed task %v\n", id, ch)
	}
}

func main() {
	taskCh := make(chan int)
	wg := &sync.WaitGroup{}
	// wgWorker := &sync.WaitGroup{}

	wg.Add(1)
	go sender(taskCh, wg)

	// wgWorker.Add(3)
	wg.Add(3)
	go worker(1, taskCh, wg)
	go worker(2, taskCh, wg)
	go worker(3, taskCh, wg)

	wg.Wait()
	// wgWorker.Wait()
	fmt.Println("All workers finished!")
}
