// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Problem 2: Buffered Channel with Results (Easy-Medium)
// // Goal: Improve Problem 1 by using buffered channels and handling results more efficiently.

// type FetchResult struct {
// 	URL   string
// 	Data  string
// 	Error error
// }

// // fetchURLWithResult simulates fetching data and returns a result struct
// func fetchURLWithResult(url string) FetchResult {
// 	// Simulate network delay (100-500ms)
// 	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
// 	time.Sleep(delay)

// 	// Simulate occasional failures (10% chance)
// 	if rand.Float32() < 0.1 {
// 		return FetchResult{
// 			URL:   url,
// 			Data:  "",
// 			Error: fmt.Errorf("network error for %s", url),
// 		}
// 	}

// 	return FetchResult{
// 		URL:   url,
// 		Data:  fmt.Sprintf("Data from %s (fetched in %v)", url, delay),
// 		Error: nil,
// 	}
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
// 		"https://medium.com",
// 		"https://dev.to",
// 	}

// 	fmt.Println("Starting buffered channel fetch...")
// 	start := time.Now()

// 	// TODO: Implement concurrent fetching using buffered channels
// 	// 1. Create a buffered channel with capacity equal to number of URLs
// 	// 2. Launch goroutines to fetch each URL and send results to channel
// 	// 3. Close the channel after all goroutines are launched
// 	// 4. Use 'for item := range channel' to collect all results
// 	// 5. Handle errors appropriately
// 	// 6. Print all results with success/failure status

// 	// Your implementation goes here:
// 	ch := make(chan FetchResult, len(urls))
// 	for _, url := range urls {
// 		go func(url string) {
// 			ch <- fetchURLWithResult(url)
// 		}(url)
// 	}

// 	// DOES NOT PRINT ALL RESULTS BECAUSE THE CHANNEL IS NOT CLOSED
// 	// close(ch)

// 	// for result := range ch {
// 	// 	fmt.Printf("URL: %s, Data: %s, Error: %v\n", result.URL, result.Data, result.Error)
// 	// }

// 	for i := 0; i < len(urls); i++ {
// 		result := <-ch
// 		fmt.Printf("URL: %s, Data: %s, Error: %v\n", result.URL, result.Data, result.Error)
// 	}
// 	close(ch)

// 	elapsed := time.Since(start)
// 	fmt.Printf("Total time: %v\n", elapsed)
// }
