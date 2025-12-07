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

## Problem 4: Fan-Out Fan-In Pattern (Medium-Hard)

### Problem Statement

Implement a fan-out/fan-in pattern:

1. **Fan-Out (Generator)**: Create a function that generates numbers 1-20 and sends them to a channel
2. **Fan-Out (Workers)**: Create **4 worker goroutines** that each:
   - Receive numbers from the input channel
   - Square each number (multiply by itself)
   - Send the squared result to an output channel
3. **Fan-In (Collector)**: Create a **collector goroutine** that:
   - Receives all squared results from the output channel
   - Collects them in a slice
   - Prints the sorted results at the end
4. Ensure proper channel closure and synchronization
5. Use `sync.WaitGroup` to coordinate all goroutines

### Expected Output
```
Squared results: [1 4 9 16 25 36 49 64 81 100 121 144 169 196 225 256 289 324 361 400]
```

### My Submission

```go
func Generator(ch chan int) {
	for i := 1; i <= 20; i++ {
		ch <- i
	}
	close(ch)
}

func Workers(ch chan int, outCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		val := c * c
		outCh <- val
	}
}

func Collector(outCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	results := []int{}
	for ch := range outCh {
		results = append(results, ch)
	}
	sort.Ints(results)
	fmt.Println("Squared results:", results)
}

func main() {
	var wg sync.WaitGroup
	var collectorWg sync.WaitGroup
	ch := make(chan int)
	outCh := make(chan int)

	collectorWg.Add(1)
	go Collector(outCh, &collectorWg)

	go Generator(ch)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Workers(ch, outCh, &wg)
	}
	wg.Wait()

	close(outCh)

	collectorWg.Wait()
}
```

### Issues Found

1. ✅ **Correct**: Proper channel closure sequence
2. ✅ **Correct**: WaitGroup usage for coordination
3. ✅ **Correct**: Execution order (collector → workers → generator)
4. ⚠️ **Minor**: Typo "Squard" → "Squared" in output

### Key Concepts Learned

- **Fan-Out**: One channel feeding multiple workers (Generator → 4 Workers)
- **Fan-In**: Multiple workers feeding one channel (4 Workers → Collector)
- **Pipeline Pattern**: Data flows through stages (generator → workers → collector)
- **Channel Closure Coordination**: 
  - Generator closes input channel after sending
  - Main closes output channel after all workers finish
  - Collector stops when output channel closes
- **Execution Order Matters**: With unbuffered channels, start receivers before senders
- **Multiple WaitGroups**: Use separate WaitGroups for different stages of the pipeline

### Important Learning Points

**Why the execution order matters:**
- Unbuffered channels block until sender/receiver are both ready
- If Generator runs before Workers, it blocks on first send
- Solution: Start Collector and Workers first, then Generator

**Channel closure sequence:**
```
1. Generator finishes → close(input channel)
2. Workers finish (detected by wg.Wait()) → close(output channel)
3. Collector stops when output channel closes
```

---

## Q&A Section

### Question: Buffered vs Unbuffered Channels - Why deadlock?

**Answer:**

**Buffered Channel:**
```go
task := make(chan int, 10)  // Can hold 10 values
for i := 1; i <= 10; i++ {
    task <- i  // Non-blocking (until buffer full)
}
// Workers can start after this - simpler execution order
```

**Unbuffered Channel:**
```go
task := make(chan int)  // Buffer size = 0
// If you send before receivers are ready → DEADLOCK
for i := 1; i <= 10; i++ {
    task <- i  // BLOCKS waiting for receiver
}
```

**Solution for unbuffered channels:**
- Start receivers (workers) BEFORE sending
- Or send in a goroutine while receivers are running

**When to Use Each:**

- **Buffered Channels**: 
  - When you know approximate max items
  - Want to decouple sender/receiver timing
  - Most practical for many real-world cases
  - Simpler code, more forgiving

- **Unbuffered Channels**:
  - When you need strict synchronization
  - Building reusable libraries
  - Production code where correctness is critical
  - More complex but "textbook correct"

---

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

## Additional Intermediate Practice Problems

### Problem 4A: Simple Producer-Consumer (Easy-Medium)

**Problem Statement:**

Create a simple producer-consumer pattern:

1. Create **2 producer goroutines** that each generate 5 random numbers (1-100) and send them to a shared channel
2. Create **1 consumer goroutine** that receives all numbers from the channel and prints them
3. Use proper channel closing to signal completion
4. Ensure all goroutines complete before the program exits

**Expected Output (format may vary):**
```
Producer 1: sent 42
Producer 2: sent 17
Consumer: received 42
Consumer: received 17
Producer 1: sent 88
...
Consumer: received 88
All done!
```

**My Initial Submission:**

```go
func producer(id int, shared chan int) {
	for i := 0; i < 5; i++ {
		val := rand.IntN(100)
		fmt.Printf("Producer %v: sent %v\n", id, val)
		shared <- val
	}
}

func main() {
	shared := make(chan int)
	go producer(1, shared)
	go producer(2, shared)
	for i := 0; i < 10; i++ {
		val := <-shared
		fmt.Printf("Consumer: received %v\n", val)
	}
	fmt.Println("All done!")
	close(shared)
}
```

**My Submission with WaitGroup (Attempt 1 - Deadlock):**

```go
func producer(id int, shared chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		val := rand.IntN(100)
		fmt.Printf("Producer %v: sent %v\n", id, val)
		shared <- val
	}
}

func main() {
	shared := make(chan int)  // Unbuffered channel
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(1, shared, &wg)
	go producer(2, shared, &wg)

	wg.Wait()      // Deadlock! Producers block on send, no receiver ready
	close(shared)

	for val := range shared {
		fmt.Printf("Consumer: received %v\n", val)
	}
	fmt.Println("All done!")
}
```

**Issues Found:**

1. ❌ **Deadlock with Unbuffered Channel**: When using unbuffered channels, the execution order matters. If producers start sending before the consumer is ready, they block forever waiting for a receiver.

2. ⚠️ **Channel Closing**: In the initial submission, channel was closed in main after receiving, which works but isn't ideal.

**Solution 1: Unbuffered Channel with Proper Order**

```go
func producer(id int, shared chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		val := rand.IntN(100)
		fmt.Printf("Producer %v: sent %v\n", id, val)
		shared <- val
	}
}

func main() {
	shared := make(chan int)  // Unbuffered
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Step 1: Start consumer FIRST (ready to receive)
	consumerWg.Add(1)
	go func() {
		defer consumerWg.Done()
		for val := range shared {
			fmt.Printf("Consumer: received %v\n", val)
		}
	}()

	// Step 2: Start producers (consumer is ready!)
	producerWg.Add(2)
	go producer(1, shared, &producerWg)
	go producer(2, shared, &producerWg)

	// Step 3: Wait for producers
	producerWg.Wait()
	close(shared)

	// Step 4: Wait for consumer
	consumerWg.Wait()
	fmt.Println("All done!")
}
```

**Solution 2: Buffered Channel (Simpler and More Practical)**

```go
func producer(id int, shared chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		val := rand.IntN(100)
		fmt.Printf("Producer %v: sent %v\n", id, val)
		shared <- val  // Non-blocking (until buffer full)
	}
}

func main() {
	shared := make(chan int, 10)  // Buffered channel
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(1, shared, &wg)
	go producer(2, shared, &wg)

	wg.Wait()      // Wait for producers
	close(shared)  // Close channel

	// Consumer receives all values
	for val := range shared {
		fmt.Printf("Consumer: received %v\n", val)
	}

	fmt.Println("All done!")
}
```

**Comparison of Approaches:**

| Approach | Pros | Cons | When to Use |
|----------|------|------|-------------|
| **Fixed Count Loop** | Simple, no WaitGroup needed | Hard-coded count, breaks if count changes | Learning, simple scripts |
| **Buffered + for range** | Flexible, simpler execution order, no hard-coded count | Need to know buffer size | Most practical for real-world cases |
| **Unbuffered + WaitGroup** | Most "correct" pattern, no buffer needed | More complex, execution order critical | Production code, strict synchronization |

**Key Concepts Learned:**

- **Unbuffered Channel Deadlock**: With unbuffered channels, receivers must be ready before senders start. Otherwise, senders block forever.
- **Buffered Channels Simplify**: Buffered channels decouple sender/receiver timing, making code simpler and more forgiving.
- **Execution Order Matters**: With unbuffered channels: Consumer → Producers → Wait → Close → Wait for consumer
- **WaitGroup Coordination**: Use separate WaitGroups for different stages (producers vs consumer)
- **for range Pattern**: Automatically receives until channel closes - no need to know exact count

**Real-World Perspective:**

In practice, many developers use buffered channels for producer-consumer patterns because:
1. Simpler code (no complex execution order)
2. Better performance (less blocking)
3. More forgiving (order matters less)

The unbuffered + WaitGroup pattern is more "textbook correct" but often overkill for simple cases. Choose based on your needs:
- **Learning/Simple**: Fixed count loop
- **Practical/Real-world**: Buffered channel + for range
- **Production/Strict**: Unbuffered + WaitGroup with proper ordering

---

### Problem 4B: Channel with Timeout (Easy-Medium)

**Problem Statement:**

Create a program that demonstrates channel timeout:

1. Create a function that sends numbers 1-5 to a channel with a 200ms delay between each
2. In main, receive from the channel with a **1 second timeout**
3. Print each received number
4. If timeout occurs, print "Timeout! Not all numbers received"

**Expected Output:**
```
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
All numbers received!
```

**Or if timeout:**
```
Received: 1
Received: 2
Received: 3
Timeout! Not all numbers received
```

**My Submission:**

```go
func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)
	go sender(ch)

	var val int
	for i := 1; i <= 5; i++ {
		select {
		case val = <-ch:
			fmt.Printf("Received: %v\n", val)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout! Not all numbers receieved")
		}
	}
	fmt.Println("All numbers received!")
}
```

**Observation:**

The timeout never triggers in this implementation because:
- Sender sends immediately, then sleeps 200ms
- Each value arrives within 200ms (well under the 1-second timeout)
- The timeout case is never reached

**What Problem 4B is Actually Teaching:**

The goal is to learn the **pattern**, not necessarily to trigger the timeout. This problem teaches:

1. **The `select` Statement Syntax**: How to use `select` with multiple cases
2. **Timeout Pattern**: How to add timeouts to channel operations using `time.After()`
3. **The "Receive with Timeout" Pattern**: A common pattern in production code

**Why the Pattern Matters (Even if Timeout Doesn't Trigger):**

Think of it like a seatbelt:
- You wear it even if you don't crash
- It's there for safety
- The timeout protects against:
  - Network delays
  - Slow operations
  - Deadlocks
  - Hung processes

In real-world code, you add timeouts as a safety mechanism, even if they rarely trigger. The pattern is what's important.

**Key Concepts Learned:**

- **`select` Statement**: Handles multiple channel operations, executes whichever case is ready first
- **`time.After()`**: Creates a channel that fires after the specified duration
- **Timeout Pattern**: Always have a timeout case in `select` for safety in production code
- **Pattern vs. Execution**: The pattern is what matters, not whether the timeout actually triggers in this specific example

**Understanding `select` with Timeout:**

```go
select {
case val = <-ch:
    // Handle received value (executes if data is ready)
case <-time.After(1 * time.Second):
    // Handle timeout (executes if 1 second passes without data)
}
```

- `select` waits for whichever case is ready first
- If channel has data → receives immediately
- If no data arrives within timeout → timeout fires
- Each `select` has its own timeout timer

**To Test Timeout (Optional):**

If you want to see the timeout trigger, add a delay in the sender:

```go
func sender(ch chan int) {
	time.Sleep(1500 * time.Millisecond)  // Wait 1.5s before first send
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
}
```

This will trigger the timeout on the first receive. However, the real learning objective is understanding the pattern, not triggering the timeout.

**Real-World Application:**

This pattern is used extensively in production code:
- API calls with timeouts
- Database queries with timeouts
- Network operations with timeouts
- Any operation that might hang or take too long

The timeout may not trigger in this simple example, but the pattern is essential for robust concurrent code.

---

### Problem 4C: Multiple Channels with Select (Medium)

**Problem Statement:**

Create a program that handles multiple channels:

1. Create **2 channels**: one for numbers (1-5) and one for letters ('a'-'e')
2. Create goroutines to send to each channel
3. Use a **single goroutine** with `select` to:
   - Print "Received number: X" when receiving from numbers channel
   - Print "Received letter: X" when receiving from letters channel
   - Exit when both channels are closed
4. Close channels after sending all data

**Expected Output (order may vary):**
```
Received number: 1
Received letter: a
Received number: 2
Received letter: b
...
Received number: 5
Received letter: e
Done!
```

**My Submission (Final Working Version):**

```go
func sendNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func sendLetter(ch chan rune, wg *sync.WaitGroup) {
	defer wg.Done()
	s := "abcde"
	for _, c := range s {
		ch <- c
	}
	close(ch)
}

func main() {
	chNum := make(chan int)
	chLetter := make(chan rune)
	wgNum := &sync.WaitGroup{}
	wgLetter := &sync.WaitGroup{}
	wgReceiver := &sync.WaitGroup{}

	wgReceiver.Add(1)
	go func() {
		numClosed := false
		letterClosed := false
		defer wgReceiver.Done()
		for {
			select {
			case num, ok := <-chNum:
				if !ok {
					numClosed = true
					chNum = nil  // Set to nil to exclude from select
				} else {
					fmt.Printf("Received number: %v\n", num)
				}
			case letter, ok := <-chLetter:
				if !ok {
					letterClosed = true
					chLetter = nil  // Set to nil to exclude from select
				} else {
					fmt.Printf("Received letter: %v\n", string(letter))
				}
			}
			if numClosed && letterClosed {
				break
			}
		}
	}()

	wgNum.Add(1)
	wgLetter.Add(1)
	go sendNum(chNum, wgNum)
	go sendLetter(chLetter, wgLetter)
	wgNum.Wait()
	wgLetter.Wait()

	wgReceiver.Wait()
	fmt.Println("Done!")
}
```

**Issues Encountered and Fixed:**

1. **Initialization Bug**: Initially used `var numOK bool` and `var letterOK bool` which default to `false`. The exit condition `if !numOK && !letterOK` was `true` from the start, causing immediate exit. **Fix**: Use separate `numClosed` and `letterClosed` flags initialized to `false`.

2. **Execution Order**: Initially tried to wait for receiver before starting senders, causing deadlock. **Fix**: Start receiver goroutine first, then start senders, then wait for senders, then wait for receiver.

3. **Channel Closure Detection**: Need to check `ok` value when receiving to detect closed channels. **Fix**: Use `value, ok := <-ch` syntax and check `!ok` to detect closure.

4. **Closed Channel Handling**: When a channel is closed, `select` can still select it repeatedly, returning `(zeroValue, false)`. **Fix**: Set closed channels to `nil` when `!ok` is detected. A `nil` channel in `select` is never selected, preventing repeated zero-value returns.

**Key Concepts Learned:**

- **`select` with Multiple Channels**: Can handle multiple channels simultaneously, executing whichever case is ready first
- **Channel Closure Detection**: Use `value, ok := <-ch` syntax. When `ok` is `false`, the channel is closed
- **Execution Order with Unbuffered Channels**: 
  - Start receiver goroutine first (so it's ready to receive)
  - Then start senders
  - Wait for senders to finish (they close channels)
  - Then wait for receiver to finish (it detects closure and exits)
- **Why Receiver Must Start First**: With unbuffered channels, senders block until a receiver is ready. Starting the receiver first maximizes the chance it's in `select` when senders start sending.
- **Why Wait for Senders Before Receiver**: Senders must finish sending and close channels before the receiver can detect closure and exit. If you wait for receiver first, it might still be waiting for data that hasn't been sent yet → deadlock.

**Understanding the Execution Order:**

```
Time →
1. Start receiver goroutine (line 35)
   ↓ (receiver enters select, waiting)
2. Start sender goroutines (lines 64-65)
   ↓ (senders send data, receiver receives)
3. Wait for senders to finish (lines 66-67)
   ↓ (senders close channels)
4. Wait for receiver to finish (line 69)
   ↓ (receiver detects both channels closed, exits)
5. "Done!"
```

**Why Setting Channels to `nil` Works:**

Setting a closed channel to `nil` in `select` excludes it from future selections. This is important because:
- A `nil` channel in `select` is **never selected** (it's always blocked)
- This prevents `select` from repeatedly selecting closed channels and returning `(zeroValue, false)`
- You must set it to `nil` when `!ok` (channel closed), not in the `else` block
- The channel variable must be captured from the outer scope (which works in this case)

**Key Point**: When a channel is closed and set to `nil`, the `select` statement will skip that case entirely, making the loop more efficient and preventing unnecessary zero-value processing.

**Real-World Application:**

This pattern is used for:
- Handling multiple data sources concurrently
- Merging data from multiple channels
- Implementing fan-in patterns
- Processing data from multiple workers simultaneously

The key is proper synchronization: ensuring receivers are ready before senders start, and waiting for senders to finish before waiting for receivers.

---

### Problem 4D: WaitGroup Practice (Easy-Medium)

**Problem Statement:**

Practice using WaitGroup properly:

1. Create **5 goroutines** that each:
   - Print "Goroutine X: starting..."
   - Sleep for a random duration (100-500ms)
   - Print "Goroutine X: done!"
2. Use `sync.WaitGroup` to wait for all goroutines
3. Print "All goroutines finished!" after all complete

**Expected Output (order may vary):**
```
Goroutine 1: starting...
Goroutine 2: starting...
Goroutine 3: starting...
Goroutine 4: starting...
Goroutine 5: starting...
Goroutine 3: done!
Goroutine 1: done!
...
Goroutine 5: done!
All goroutines finished!
```

**Hints:**
- Call `wg.Add(5)` before starting goroutines
- Each goroutine should call `defer wg.Done()`
- Main calls `wg.Wait()` to wait

**Key Learning:**
- Proper WaitGroup usage pattern
- `defer wg.Done()` ensures Done is always called
- WaitGroup coordinates multiple independent goroutines

---

### Problem 4E: Buffered Channel Practice (Easy-Medium)

**Problem Statement:**

Understand buffered channels:

1. Create a **buffered channel** with capacity 3
2. Send 5 numbers (1-5) to the channel
3. Create a receiver that prints all numbers
4. Observe the behavior difference between buffered and unbuffered

**Expected Output:**
```
Sending 1 (non-blocking, buffer has space)
Sending 2 (non-blocking, buffer has space)
Sending 3 (non-blocking, buffer has space)
Sending 4 (blocks until receiver takes one)
Sending 5 (blocks until receiver takes one)
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
```

**Hints:**
- Buffered channels allow sending up to buffer size without blocking
- After buffer is full, sends block until space is available
- Receivers can drain the buffer

**Key Learning:**
- Buffered vs unbuffered channel behavior
- When sends block with buffered channels
- Decoupling sender/receiver timing

---

### Problem 4F: Channel Direction (Easy)

**Problem Statement:**

Practice using channel directions (send-only, receive-only):

1. Create a function `sender(ch chan<- int)` that sends numbers 1-5 (send-only channel)
2. Create a function `receiver(ch <-chan int)` that receives and prints (receive-only channel)
3. In main, create channel and pass to both functions
4. Use proper channel closure

**Expected Output:**
```
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
All done!
```

**Hints:**
- `chan<- int` means send-only
- `<-chan int` means receive-only
- This provides type safety and makes code intent clear

**Key Learning:**
- Channel direction types for better code safety
- Enforcing send-only or receive-only at compile time
- Clearer function signatures

---

## Summary of Common Mistakes

1. **Channel Closing**: Always close channels in the sender, not the receiver
2. **Closure Bug**: When using loop variables in goroutines, pass them as parameters instead of capturing them in closures
3. **WaitGroup Usage**: Remember to call `wg.Done()` (usually with `defer`) in each goroutine
4. **Channel Range**: Use `for range ch` to automatically handle channel closure
5. **Random Range**: Use `rand.Intn(max - min + 1) + min` for inclusive ranges
6. **Unbuffered Channel Deadlock**: With unbuffered channels, receivers must be ready before senders start. Otherwise, senders block forever. Solution: Start receivers first, or use buffered channels.
7. **Execution Order with Unbuffered Channels**: When using unbuffered channels, the order matters - start receivers before senders, or use buffered channels for simpler code.

---

## Recommended Learning Path

### Beginner → Easy (Completed ✅)
- ✅ Problem 1: Basic Goroutine Execution
- ✅ Problem 2: Basic Channel Communication

### Easy → Medium (Completed ✅)
- ✅ Problem 3: Worker Pool Pattern

### Medium → Medium-Hard (Completed ✅)
- ✅ Problem 4: Fan-Out Fan-In Pattern

### Intermediate Practice (Recommended Next)
- Problem 4A: Simple Producer-Consumer (Easy-Medium)
- Problem 4B: Channel with Timeout (Easy-Medium)
- Problem 4C: Multiple Channels with Select (Medium)
- Problem 4D: WaitGroup Practice (Easy-Medium)
- Problem 4E: Buffered Channel Practice (Easy-Medium)
- Problem 4F: Channel Direction (Easy)

### Advanced Problems (After mastering intermediate)
- Problem 5: Context Cancellation & Timeout (Medium-Hard)
- Problem 6: Select Statement with Multiple Channels (Hard)
- Problem 7: Mutex for Shared State (Hard)
- Problem 8: Pipeline Pattern with Error Handling (Very Hard)

---

## Learning Tips

1. **Master the basics first**: Make sure you're comfortable with Problems 1-3 before moving on
2. **Practice intermediate problems**: Problems 4A-4F build essential skills gradually
3. **Understand the patterns**: Each problem teaches a specific pattern you'll use in real code
4. **Common gotchas to watch for**:
   - Closure bugs with loop variables
   - Channel closure timing
   - Unbuffered channel deadlocks
   - WaitGroup usage (Add before goroutines, Done in goroutines)
5. **Test your code**: Run it multiple times - concurrency bugs can be non-deterministic

