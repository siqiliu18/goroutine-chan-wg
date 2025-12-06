package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(workerID int, taskID chan int) {
	delay := rand.Intn(401) + 100
	for id := range taskID {
		time.Sleep(time.Millisecond * time.Duration(delay))
		fmt.Printf("Worker %v processed task %v\n", workerID, id)
	}
}

func main() {
	var wg sync.WaitGroup
	task := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		task <- i
	}
	close(task)

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func() {
			defer wg.Done()
			worker(i, task)
		}()
	}
	wg.Wait()

	fmt.Println("All workers finished!")
}
