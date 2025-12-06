package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Problem 1: Basic Concurrent Fetching (Easy)
// Goal: Fetch data from multiple URLs concurrently using unbuffered channels.

// fetchURL simulates fetching data from a URL
// It takes a random amount of time to simulate network delay
func fetchURL(url string) string {
	// Simulate network delay (100-500ms)
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)

	// Simulate some data being fetched
	return fmt.Sprintf("Data from %s (fetched in %v)", url, delay)
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// URLs to fetch
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	fmt.Println("Starting concurrent fetch...")
	start := time.Now()

	// TODO: Implement concurrent fetching using unbuffered channels
	// 1. Create an unbuffered channel to collect results
	// 2. Launch goroutines to fetch each URL
	// 3. Collect all results from the channel
	// 4. Print all results

	// Your implementation goes here:
	// Sender blocks until receiver is ready
	// Goroutine can't proceed until main goroutine reads from channel
	// Forces synchronization: send → receive → send → receive
	// Slower because of blocking
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ch <- fetchURL(url)
		}(url)
	}
	/*
		ch := make(chan string)
		wg := sync.WaitGroup{}

		for _, url := range urls {
		    wg.Add(1)  // ← Increment counter for each goroutine
		    go func(url string) {
		        defer wg.Done()  // ← Decrement when goroutine finishes
		        ch <- fetchURL(url)
		    }(url)
		}

		// Close the channel after all goroutines finish
		go func() {
		    wg.Wait()  // ← Wait until all goroutines call Done()
		    close(ch)
		}()

		// Now this loop will exit when ch is closed
		for msg := range ch {
		    fmt.Println("Received:", msg)
		}
	*/

	// This will never reach because the ch is never closed
	// the loop blocks forever waiting
	// for msg := range ch {
	// 	fmt.Println("Received:", msg)
	// }
	// WHY CANNOT close(ch) here, because ← Might close while goroutines are still sending!

	// Loop exits after exactly 5 iterations
	// Channel is never closed, but we're not reading from it anymore
	// You explicitly receive exactly N messages (5 in this case), so the loop exits by itself.
	// There's no for...range waiting indefinitely — you know exactly when to stop.
	// The channel just becomes unused afterward, which is fine.
	for i := 0; i < len(urls); i++ {
		fmt.Println("Received:", <-ch)
	}
	// Why it's safe: By the time you exit the loop, all 5 goroutines have already sent their data and finished.
	// The channel is no longer being written to, so closing it is safe.
	close(ch) // better to have after this for i < len(urls)

	elapsed := time.Since(start)
	fmt.Printf("Total time: %v\n", elapsed)

	/*
		Summary:

			Approach 2: You know exactly when all data arrives (after 5 reads), so close is safe
			Approach 1: You don't know when goroutines finish, so you must wait with wg.Wait() before closing
	*/
}
