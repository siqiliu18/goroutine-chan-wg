package main

import (
	"fmt"
	"sync"
)

func myGoroutine() {
	fmt.Println("This is my first goroutine")
}

func myGoroutine2() {
	fmt.Println("This is my second goroutine")
}

func myGoroutine3(c chan<- struct{}) {
	fmt.Println("This is my third goroutine")
	c <- struct{}{}
}

func myGoroutine4() {
	fmt.Println("This is my fourth goroutine")
}

func myGoroutine5(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("This is my fifth goroutine")
}

func main() {
	/* What is the problem here?
		go myGoroutine()
	} // closing the main block
	*/

	/* https://stackoverflow.com/a/18215175/9954367
	WaitGroups are definitely the canonical way to do this.
	Just for the sake of completeness, though, here's the solution that was commonly used before WaitGroups were introduced.
	The basic idea is to use a channel to say "I'm done," and have the main goroutine wait until each spawned routine has reported its completion.
	*/
	// method 1:
	c1 := make(chan struct{})
	go func() {
		myGoroutine()
		myGoroutine2()
		c1 <- struct{}{}
	}()
	<-c1

	// method 2 - set chan as a param of the func and pass it an chann obj.
	c2 := make(chan struct{})
	go myGoroutine3(c2)
	<-c2 // make sure to std out the chann to tell the main thread to wait.

	// wait group - https://www.geeksforgeeks.org/using-waitgroup-in-golang/
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		myGoroutine4()
	}()

	wg.Add(1)
	myGoroutine5(wg)

	wg.Wait()
}
