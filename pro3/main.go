package main

import "fmt"

// https://stackoverflow.com/questions/18660533/why-does-the-use-of-an-unbuffered-channel-in-the-same-goroutine-result-in-a-dead
func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

/*
1. Based on unbuffered chann's definition - the sender blocks until the receiver has received the value.
	- when a channel is full, the sender waits for another goroutine to make some room by receiving
	- you can see an unbuffered channel as an always full one : there must be another goroutine to take what the sender sends.
	- c <- 1 : blocks because the channel is unbuffered. As there's no other goroutine to receive the value, the situation can't resolve, this is a deadlock.
2. You can make it not blocking by changing the channel creation to - c := make(chan int, 1)
	- so that there's room for one item in the channel before it blocks.
3. But that's not what concurrency is about. Normally, you wouldn't use a channel without other goroutines to handle what you put inside. You could define a receiving goroutine like this:
	func main() {
		c := make(chan int)
		go func() {
			fmt.Println("received:", <-c)
		}()
		c <- 1
	}
4. If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.
	- If the amount of receivers are bigger than the buffer size, then deadlock could still happen
   ch := make(chan int, 1) // channel Buffer size 1
   ch <- 10
   ch <- 20 // Blocked: Because Buffer size is already full and no one is waiting to recieve the Data  from channel
   <- ch
   <- ch
*/
