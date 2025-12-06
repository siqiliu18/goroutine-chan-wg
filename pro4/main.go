package main

import (
	"fmt"
	"time"
)

func fast(num int, out chan<- int) {
	result := num * 2
	time.Sleep(5 * time.Millisecond)
	out <- result
}

func slow(num int, out chan<- int) {
	result := num * 2
	time.Sleep(15 * time.Millisecond)
	out <- result
}

func main() {
	channels := make([]chan int, 2)
	channels[0] = make(chan int)
	channels[1] = make(chan int)
	/*
		OR:
			var channels []chan int
			channels = append(channels, make(chan int))
			channels = append(channels, make(chan int))
	*/

	go fast(2, channels[0])
	go slow(3, channels[1])

	for i := 0; i < 2; i++ {
		select {
		case res := <-channels[0]:
			fmt.Println("fast finished first, result: ", res)
		case res := <-channels[1]:
			fmt.Println("slow finished first, result: ", res)
		}
	}
}
