# Goroutine Coding Interview Practice

This document contains the problems, my submissions, issues found, and Q&A from goroutine practice sessions.

---

## Problem 1: Basic Goroutine Execution (Beginner)

### Problem Statement

Write a program that demonstrates basic goroutine execution:

1. Create a function `sayHello(name string)` that prints "Hello, [name]!" **3 times** with a small delay between each print
2. Call this function as a **goroutine** for 3 different names: "Alice", "Bob", and "Charlie"
3. Use `time.Sleep()` in the main function to wait for all goroutines to finish
4. Print "All greetings done!" at the end

### Expected Output
```
Hello, Alice!
Hello, Bob!
Hello, Charlie!
Hello, Alice!
Hello, Bob!
Hello, Charlie!
Hello, Alice!
Hello, Bob!
Hello, Charlie!
All greetings done!
```

### My Initial Submission

```go
func sayHello(name string) {
	delay := rand.Intn(1000)
	for i := 0; i < 3; i++ {
		sleep := time.Duration(delay) * time.Millisecond
		time.Sleep(sleep)
		fmt.Printf("Hello, [%v]!", name)
		fmt.Println()
	}
}

func main() {
	names := []string{"Alics", "Bob", "Charlie"}
	for _, name := range names {
		go sayHello(name)
	}
	time.Sleep(5 * time.Second)
}
```

### Issues Found

1. ✅ **Fixed**: Added loop to print 3 times per name
2. ✅ **Fixed**: Added "All greetings done!" message at the end
3. ✅ **Fixed**: Typo "Alics" → "Alice"

### Key Concepts Learned

- Goroutines run concurrently (order may vary)
- Main function needs to wait, otherwise it exits before goroutines finish
- Each goroutine executes independently
- Using `go` keyword launches a function as a goroutine

---

## Problem 2: Basic Channel Communication (Easy)

### Problem Statement

Write a program that uses channels to pass data between goroutines:

1. Create a function `sendNumbers(ch chan int)` that sends numbers 1, 2, 3, 4, 5 to a channel
2. Create a function `receiveNumbers(ch chan int)` that receives numbers from the channel and prints them
3. In main:
   - Create a channel
   - Launch `sendNumbers` as a goroutine
   - Call `receiveNumbers` in main (not as goroutine)
   - Close the channel after sending is done

### Expected Output
```
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
All numbers received!
```

### My Submission

```go
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func receiveNumbers(ch chan int) {
	fmt.Println("Received", <-ch)
}

func main() {
	ch := make(chan int)
	go sendNumbers(ch)
	for i := 1; i <= 5; i++ {
		receiveNumbers(ch)
	}
	fmt.Println("All numbers received!")
	close(ch)
}
```

### Issues Found

1. ❌ **Critical Issue**: Channel closing - The channel should be closed by the **sender** (`sendNumbers`), not by the receiver (main). In Go, the sender is responsible for closing the channel.

2. ⚠️ **Minor**: Output format - Code prints "Received 1" but expected format was "Received: 1" (with colon)

### Corrected Version

```go
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)  // Close channel after sending all numbers
}

func receiveNumbers(ch chan int) {
	for num := range ch {  // This loop automatically stops when channel is closed
		fmt.Printf("Received: %d\n", num)
	}
}

func main() {
	ch := make(chan int)
	go sendNumbers(ch)
	receiveNumbers(ch)  // No need for loop here - receiveNumbers handles it
	fmt.Println("All numbers received!")
}
```

### Key Concepts Learned

- **Sender closes the channel** - This is a Go best practice
- `for range ch` automatically stops when the channel is closed
- Channels are the primary way to communicate between goroutines
- `ch <- value` to send, `value := <-ch` to receive

---

## Problem 3: Worker Pool Pattern (Medium)

### Problem Statement

Implement a worker pool that processes tasks concurrently:

1. Create a **task channel** that receives integers (task IDs: 1-10)
2. Create **3 worker goroutines** that process tasks from the channel
3. Each worker should:
   - Receive a task ID from the channel
   - Simulate work by sleeping for a random duration (100-500ms)
   - Print "Worker X processed task Y"
4. Send all 10 tasks (1-10) to the channel
5. Close the channel after sending all tasks
6. Wait for all workers to finish before printing "All workers finished!"

### Expected Output (order may vary)
```
Worker 1 processed task 1
Worker 2 processed task 2
Worker 3 processed task 3
Worker 1 processed task 4
Worker 2 processed task 5
...
Worker 3 processed task 10
All workers finished!
```

### My Submission

```go
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
```

### Issues Found

1. ❌ **Critical Bug - Closure Issue**: On lines 28-31, capturing `i` in the goroutine closure. All goroutines may see the same value of `i` (likely 4 after the loop), causing incorrect worker IDs.

2. ⚠️ **Minor**: Delay calculation - The delay is calculated once per worker (line 11), so all tasks from the same worker use the same delay. Consider calculating it per task for more realistic simulation.

### Corrected Version

```go
func worker(workerID int, taskID chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for id := range taskID {
		delay := rand.Intn(401) + 100 // 100-500ms range, calculated per task
		time.Sleep(time.Millisecond * time.Duration(delay))
		fmt.Printf("Worker %v processed task %v\n", workerID, id)
	}
}

func main() {
	var wg sync.WaitGroup
	task := make(chan int, 10)
	
	// Send all tasks
	for i := 1; i <= 10; i++ {
		task <- i
	}
	close(task)

	// Launch 3 workers
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, task, &wg) // Pass i directly, not in closure
	}
	
	wg.Wait()
	fmt.Println("All workers finished!")
}
```

### Key Concepts Learned

- **Closure Bug**: When using loop variables (`i`) inside goroutines, all goroutines may reference the same variable. Always pass loop variables as parameters.
- **Worker Pool Pattern**: Multiple workers processing from a shared task queue
- `sync.WaitGroup`: Coordinate waiting for multiple goroutines
- Channel closure signals workers to stop (via `for range`)

---

## Q&A Section

### Question: How to calculate `rand.Intn(401) + 100` for 100-500ms range?

**Answer:**

#### How `rand.Intn()` works:
- `rand.Intn(n)` returns a random integer in the range **[0, n)**, meaning 0 to n-1 (inclusive of 0, exclusive of n)
- Examples:
  - `rand.Intn(5)` returns: 0, 1, 2, 3, or 4
  - `rand.Intn(10)` returns: 0, 1, 2, ..., 9

#### To get a range from `min` to `max` (inclusive):

Use the formula: **`rand.Intn(max - min + 1) + min`**

#### For 100-500ms:
- `min = 100`
- `max = 500`
- `rand.Intn(500 - 100 + 1) + 100`
- `rand.Intn(401) + 100`

#### Step-by-step:
1. `rand.Intn(401)` returns: 0, 1, 2, ..., 400 (401 possible values)
2. Adding 100 shifts the range: 100, 101, 102, ..., 500
3. Result: **100 to 500 (inclusive)**

#### Visual breakdown:
```
rand.Intn(401) gives:    0,   1,   2, ..., 400
Adding 100 gives:      100, 101, 102, ..., 500
```

#### General formula examples:
- Range 1-10: `rand.Intn(10) + 1` → `rand.Intn(10)` gives 0-9, +1 gives 1-10
- Range 50-100: `rand.Intn(51) + 50` → `rand.Intn(51)` gives 0-50, +50 gives 50-100
- Range 100-500: `rand.Intn(401) + 100` → `rand.Intn(401)` gives 0-400, +100 gives 100-500

#### Why `max - min + 1`?
- The `+1` ensures the maximum value is included
- Without it, you'd get `max - min` possible values, missing the top value

#### Quick reference:
```
Want range [a, b] (inclusive)?
→ rand.Intn(b - a + 1) + a
```

---

## Summary of Common Mistakes

1. **Channel Closing**: Always close channels in the sender, not the receiver
2. **Closure Bug**: When using loop variables in goroutines, pass them as parameters instead of capturing them in closures
3. **WaitGroup Usage**: Remember to call `wg.Done()` (usually with `defer`) in each goroutine
4. **Channel Range**: Use `for range ch` to automatically handle channel closure
5. **Random Range**: Use `rand.Intn(max - min + 1) + min` for inclusive ranges

---

## Next Problems to Practice

- Problem 4: Fan-Out Fan-In Pattern (Medium-Hard)
- Problem 5: Context Cancellation & Timeout (Medium-Hard)
- Problem 6: Select Statement with Multiple Channels (Hard)
- Problem 7: Mutex for Shared State (Hard)
- Problem 8: Pipeline Pattern with Error Handling (Very Hard)

