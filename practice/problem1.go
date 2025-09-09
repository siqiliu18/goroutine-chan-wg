// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Problem 1: Basic Concurrent Fetching (Easy)
// // Goal: Fetch data from multiple URLs concurrently using unbuffered channels.

// // fetchURL simulates fetching data from a URL
// // It takes a random amount of time to simulate network delay
// func fetchURL(url string) string {
// 	// Simulate network delay (100-500ms)
// 	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
// 	time.Sleep(delay)

// 	// Simulate some data being fetched
// 	return fmt.Sprintf("Data from %s (fetched in %v)", url, delay)
// }

// func main() {
// 	// Seed random number generator
// 	rand.Seed(time.Now().UnixNano())

// 	// URLs to fetch
// 	urls := []string{
// 		"https://example.com",
// 		"https://google.com",
// 		"https://github.com",
// 		"https://stackoverflow.com",
// 		"https://golang.org",
// 	}

// 	fmt.Println("Starting concurrent fetch...")
// 	start := time.Now()

// 	// TODO: Implement concurrent fetching using unbuffered channels
// 	// 1. Create an unbuffered channel to collect results
// 	// 2. Launch goroutines to fetch each URL
// 	// 3. Collect all results from the channel
// 	// 4. Print all results

// 	// Your implementation goes here:
// 	ch := make(chan string)
// 	for _, url := range urls {
// 		go func(url string) {
// 			ch <- fetchURL(url)
// 		}(url)
// 	}

// 	// fmt.Println("Main: About to close channel")
// 	// close(ch) // This should cause panic!
// 	// fmt.Println("Main: Closed channel")

// 	// This will never reach because close() panics
// 	// for msg := range ch {
// 	// 	fmt.Println("Received:", msg)
// 	// }
// 	for i := 0; i < len(urls); i++ {
// 		fmt.Println("Received:", <-ch)
// 	}
// 	close(ch)

// 	// close(ch) // ❌ This happens after reading, which is fine but not optimal
// 	elapsed := time.Since(start)
// 	fmt.Printf("Total time: %v\n", elapsed)
// }
