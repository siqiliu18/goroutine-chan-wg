package main

import (
	"fmt"
	"sync"
	"time"
)

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		count("sheep")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count("fish")
	}()

	wg.Wait()
}