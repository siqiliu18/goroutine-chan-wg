# Goroutine Practice Problems

## Problem: Web Scraper with Rate Limiting

You're building a web scraper that needs to fetch data from multiple URLs concurrently while respecting rate limits and handling various edge cases. This problem set covers all major goroutine concepts with increasing difficulty.

## Setup
Create a new Go module in this directory:
```bash
go mod init practice
```

## Concepts Covered
- ✅ Unbuffered channels
- ✅ Buffered channels  
- ✅ WaitGroup
- ✅ `for item := range chan`
- ✅ `select` statements
- ✅ Channel closing
- ✅ Goroutine synchronization
- ✅ Error handling in concurrent code
- ✅ Context cancellation
- ✅ Worker pools
- ✅ Fan-in/Fan-out patterns

---

## Problem 1: Basic Concurrent Fetching (Easy)
**Goal**: Fetch data from multiple URLs concurrently using unbuffered channels.

**Requirements**:
- Create a function `fetchURL(url string) string` that simulates fetching data (use `time.Sleep` to simulate network delay)
- Create 5 URLs to fetch
- Use goroutines to fetch all URLs concurrently
- Collect results using an unbuffered channel
- Print all results

**Key Concepts**: Unbuffered channels, basic goroutines

---

## Problem 2: Buffered Channel with Results (Easy-Medium)
**Goal**: Improve Problem 1 by using buffered channels and handling results more efficiently.

**Requirements**:
- Use a buffered channel with capacity equal to number of URLs
- Modify the fetching to return both URL and result
- Use `for item := range channel` to collect all results
- Close the channel properly after all goroutines complete

**Key Concepts**: Buffered channels, `for range`, channel closing

---

## Problem 3: WaitGroup Synchronization (Medium)
**Goal**: Use WaitGroup to ensure all goroutines complete before proceeding.

**Requirements**:
- Use `sync.WaitGroup` instead of relying on channel capacity
- Create a worker function that takes URL and WaitGroup as parameters
- Ensure main goroutine waits for all workers to complete
- Handle potential errors in fetching

**Key Concepts**: WaitGroup, error handling, worker functions

---

## Problem 4: Rate Limiting with Select (Medium)
**Goal**: Implement rate limiting using `select` statements and time-based channels.

**Requirements**:
- Limit to maximum 2 requests per second
- Use `time.Ticker` to create a rate limiter
- Use `select` to either send request or wait for rate limit
- Handle the case where rate limit is reached

**Key Concepts**: `select` statements, `time.Ticker`, rate limiting

---

## Problem 5: Worker Pool Pattern (Medium-Hard)
**Goal**: Implement a worker pool to process URLs with a fixed number of workers.

**Requirements**:
- Create a pool of 3 workers
- Use a job channel to distribute URLs to workers
- Use a results channel to collect results
- Workers should process jobs until the job channel is closed
- Handle graceful shutdown of workers

**Key Concepts**: Worker pools, job distribution, graceful shutdown

---

## Problem 6: Fan-in/Fan-out with Error Handling (Hard)
**Goal**: Implement fan-out (distribute work) and fan-in (collect results) patterns with comprehensive error handling.

**Requirements**:
- Fan-out: Distribute URLs to multiple worker goroutines
- Fan-in: Collect results from all workers into a single channel
- Handle errors gracefully (some URLs might fail)
- Use `select` with `default` case for non-blocking operations
- Implement timeout for individual requests

**Key Concepts**: Fan-in/Fan-out, error handling, timeouts, non-blocking operations

---

## Problem 7: Context Cancellation (Hard)
**Goal**: Implement proper cancellation using context to stop all operations when needed.

**Requirements**:
- Use `context.Context` for cancellation
- Implement a timeout for the entire operation
- Allow manual cancellation via a signal
- Ensure all goroutines respect the context cancellation
- Clean up resources properly

**Key Concepts**: Context cancellation, timeouts, resource cleanup

---

## Problem 8: Advanced Select with Multiple Channels (Expert)
**Goal**: Handle multiple types of events using advanced select patterns.

**Requirements**:
- Handle URL fetching, error reporting, and progress updates
- Use separate channels for different types of events
- Implement a coordinator that processes all event types
- Use `select` with multiple cases and proper handling
- Implement graceful shutdown with cleanup

**Key Concepts**: Multiple channel types, event coordination, advanced select patterns

---

## Bonus Challenges

### Challenge 1: Pipeline Pattern
Create a pipeline where URLs go through multiple processing stages (fetch → parse → validate → store).

### Challenge 2: Circuit Breaker
Implement a circuit breaker pattern that stops making requests if too many fail.

### Challenge 3: Load Balancing
Distribute work across multiple "servers" with different response times and capacities.

---

## Tips for Interview Success

1. **Always close channels** when done sending
2. **Use WaitGroup** for synchronization when you need to wait for goroutines
3. **Handle errors** in concurrent code - don't ignore them
4. **Use context** for cancellation and timeouts
5. **Avoid goroutine leaks** - ensure all goroutines can exit
6. **Test with edge cases** - empty input, network failures, timeouts
7. **Understand channel blocking behavior** - buffered vs unbuffered
8. **Use select for non-blocking operations** when appropriate

## Getting Started
1. Create your Go module: `go mod init practice`
2. Start with Problem 1 and work your way up
3. For each problem, create a separate file (e.g., `problem1.go`, `problem2.go`)
4. Test your solutions thoroughly
5. Ask for help when you get stuck!

Good luck! 🚀
