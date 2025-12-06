// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Problem 4: Rate Limiting with Select (Medium)
// // Goal: Implement rate limiting using `select` statements and time-based channels.

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

// 	// Simulate occasional failures (20% chance)
// 	if rand.Float32() < 0.2 {
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
// 		"https://reddit.com",
// 		"https://news.ycombinator.com",
// 		"https://dev.to",
// 	}

// 	fmt.Println("Starting rate-limited fetch...")
// 	start := time.Now()

// 	// TODO: Implement rate limiting using select statements
// 	// 1. Create a rate limiter using time.Ticker (2 requests per second)
// 	// 2. Create channels for results and rate limiting
// 	// 3. Use select to either send a request or wait for rate limit
// 	// 4. Handle the case where rate limit is reached
// 	// 5. Collect and print all results
// 	// 6. Handle errors appropriately

// 	// Your implementation goes here:
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	defer ticker.Stop()
// 	ch := make(chan FetchResult)

// 	/*
// 		What happens:

// 			<-ticker.C blocks and waits for the ticker to tick
// 			Every 500ms, the ticker sends a signal
// 			When received, launch one goroutine
// 			Wait another 500ms, launch the next one
// 			Timeline:

// 			t=0ms: Launch goroutine 1
// 			t=500ms: Launch goroutine 2
// 			t=1000ms: Launch goroutine 3
// 			t=1500ms: Launch goroutine 4
// 			... and so on
// 			So with 10 URLs and 500ms intervals = ~5 seconds minimum to launch all goroutines (plus fetch time).

// 			This is rate limiting — you control how fast goroutines are created to avoid overwhelming the system.
// 			Different from Problem 2 where all goroutines launch immediately.
// 	*/
// 	for _, url := range urls {
// 		<-ticker.C
// 		go func(url string) {
// 			ch <- fetchURLWithResult(url)
// 		}(url)
// 	}

// 	// easier way to do it is to use a for loop to receive from the channel.
// 	for i := 0; i < len(urls); i++ {
// 		res := <-ch
// 		fmt.Printf("URL: %s, Data: %s, Error: %v\n", res.URL, res.Data, res.Error)
// 	}

// 	// another way to do it is to close the channel after the range is done.
// 	/*
// 		go func() {
// 			waitGroup.Wait()
// 			close(ch)
// 		}()
// 		for data := range ch {
// 			fmt.Printf("URL: %s, Data: %s, Error: %v\n", data.URL, data.Data, data.Error)
// 		}
// 	*/

// 	elapsed := time.Since(start)
// 	fmt.Printf("Total time: %v\n", elapsed)
// }
