// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Problem 5: Worker Pool Pattern (Medium-Hard)
// // Goal: Implement a worker pool to process URLs with a fixed number of workers.

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

// 	// Simulate occasional failures (25% chance)
// 	if rand.Float32() < 0.25 {
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

// func worker(jobs <-chan string, results chan<- FetchResult) {
// 	for url := range jobs {
// 		results <- fetchURLWithResult(url)
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
// 		"https://reddit.com",
// 		"https://news.ycombinator.com",
// 		"https://dev.to",
// 		"https://stackoverflow.com",
// 		"https://golang.org",
// 		"https://medium.com",
// 		"https://dev.to",
// 		"https://reddit.com",
// 	}

// 	fmt.Println("Starting worker pool fetch...")
// 	start := time.Now()

// 	// TODO: Implement worker pool pattern
// 	// 1. Create a pool of 3 workers
// 	// 2. Use a job channel to distribute URLs to workers
// 	// 3. Use a results channel to collect results
// 	// 4. Workers should process jobs until the job channel is closed
// 	// 5. Handle graceful shutdown of workers
// 	// 6. Collect and print all results
// 	// 7. Handle errors appropriately

// 	// Your implementation goes here:

// 	/*

// 		While using unbuffered channels (jobs and results), the producer for loop (line90~92) trys to send 15 urls to channel jobs sequentially.
// 		However, there are only 3 go routines for worker, meaning after the first 3 urls sent, all 3 go routines are blocked as the main go routines is still in the loop and has not reached to the consumer code (line 95~98) yet.
// 		While using buffered channels, the producer for loop would be able to send all 15 urls to 3 channels as they have buffers with same capacity as the length of urls.
// 		And line 93 close(jobs) won't cause any deadlock as we've done sending so it is OK to close that jobs channel, right? Then consumer consumes the messages without any problem.

// 	*/
// 	jobs := make(chan string, len(urls))
// 	results := make(chan FetchResult, len(urls))

// 	for i := 0; i < 3; i++ {
// 		go worker(jobs, results)
// 	}

// 	for _, url := range urls {
// 		jobs <- url
// 	}
// 	close(jobs)

// 	for i := 0; i < len(urls); i++ {
// 		result := <-results
// 		fmt.Printf("URL: %s, Data: %s, Error: %v\n", result.URL, result.Data, result.Error)
// 	}

// 	close(results)

// 	elapsed := time.Since(start)
// 	fmt.Printf("Total time: %v\n", elapsed)
// }
