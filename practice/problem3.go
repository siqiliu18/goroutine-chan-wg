// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// // Problem 3: WaitGroup Synchronization (Medium)
// // Goal: Use WaitGroup to ensure all goroutines complete before proceeding.

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

// 	// Simulate occasional failures (15% chance)
// 	if rand.Float32() < 0.15 {
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

// // worker fetches a URL and sends result to the results channel
// // It decrements the WaitGroup when done
// func worker(url string, results chan<- FetchResult, wg *sync.WaitGroup) {
// 	defer wg.Done() // Ensure WaitGroup is decremented when function exits

// 	fmt.Printf("Worker starting to fetch: %s\n", url)
// 	result := fetchURLWithResult(url)
// 	results <- result
// 	fmt.Printf("Worker completed: %s\n", url)
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
// 		"https://reddit.com",
// 	}

// 	fmt.Println("Starting WaitGroup-based fetch...")
// 	start := time.Now()

// 	// TODO: Implement concurrent fetching using WaitGroup
// 	// 1. Create a WaitGroup
// 	// 2. Create a results channel (buffered or unbuffered - your choice)
// 	// 3. Launch worker goroutines for each URL, passing the WaitGroup
// 	// 4. Wait for all workers to complete using WaitGroup.Wait()
// 	// 5. Close the results channel
// 	// 6. Collect and print all results
// 	// 7. Handle errors appropriately
// 	wg := new(sync.WaitGroup)
// 	results := make(chan FetchResult, len(urls))
// 	for _, url := range urls {
// 		wg.Add(1)
// 		go worker(url, results, wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Printf("URL: %s, Data: %s, Error: %v\n", result.URL, result.Data, result.Error)
// 	}

// 	// Your implementation goes here:

// 	elapsed := time.Since(start)
// 	fmt.Printf("Total time: %v\n", elapsed)
// }
