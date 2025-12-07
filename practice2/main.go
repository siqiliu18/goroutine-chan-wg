package main

import (
	"fmt"
	"sync"
)

// Problem 4C

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
					chNum = nil // better to set to nil to exclude from select
				} else {
					fmt.Printf("Received number: %v\n", num)
				}
			case letter, ok := <-chLetter:
				if !ok {
					letterClosed = true
					chLetter = nil // better to set to nil to exclude from select
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

	wgReceiver.Wait() // receiver wait group must be after the senders wait groups?

	fmt.Println("Done!")
}
