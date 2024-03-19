package channel

import (
	"fmt"
	"strconv"
)

// https://medium.com/nerd-for-tech/learning-go-concurrency-goroutines-channels-8836b3c34152
// https://go101.org/article/channel.html
// https://www.atatus.com/blog/go-channels-overview/#:~:text=A%20Go%20channel%20is%20a,concurrency%20in%20apps%20and%20notifications.
func BufferChan() {
	c := make(chan int, 2) // a buffered channel
	c <- 3
	c <- 5
	// c <- 7
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	// close(c)
}

func MyGoroutine() {
	intChann := make(chan int, 3)
	for i := 0; i <= 9; i++ {
		go func(i int) {
			intChann <- i
		}(i)
	}

	for j := 0; j < 10; j++ {
		num, ok := <-intChann
		fmt.Printf("%v:%v ", num, ok)
		fmt.Println(len(intChann), cap(intChann))
	}

	close(intChann)
}

func intChannFunc(num int, intChann chan<- int) {
	intChann <- num * 2
}

func strChannFunc(str string, strChann chan<- string) {
	for i := 1; i <= 10; i++ {
		str += strconv.Itoa(i)
	}
	strChann <- str
}

func DiffTypeChann() {
	intChann := make(chan int)
	strChann := make(chan string)
	go intChannFunc(3, intChann)
	go strChannFunc("0", strChann)

	intVal := 0
	strVal := ""
	for i := 0; i < 2; i++ {
		select {
		case intVal = <-intChann:
			fmt.Println(intVal)
		case strVal = <-strChann:
			fmt.Println(strVal)
		}
	}
	close(intChann)
	close(strChann)
}
