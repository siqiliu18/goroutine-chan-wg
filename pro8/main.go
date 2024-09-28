package main

import (
	"fmt"
	"time"
)

// must watch again https://www.youtube.com/watch?v=LvgVSSpwND8 !!

func count(thing string, ch chan<-string) {
	for i := 1; i <= 5; i++ {
		// fmt.Println(i, thing)
		ch<- thing
		time.Sleep(time.Millisecond * 500)
	}
	// ch<- 0

	close(ch)
}

func main() {
	ch := make(chan string)

	go count("sheep", ch)
	// go count("fish", ch)
	
	// for {
	// 	msg, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range ch {
		fmt.Println(msg)
	}

}