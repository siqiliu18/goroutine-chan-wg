package waitgroup

import (
	"fmt"
	"sync"
)

// https://www.geeksforgeeks.org/using-waitgroup-in-golang/ (!!!!)
func runner1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am first runner")
}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func Execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go runner1(wg)
	go runner2(wg)
	wg.Wait()
}

func addOne(num int) int {
	num += 1
	return num
}

func appendStr(str string) string {
	str += "!"
	return str
}

func DiffTypeWg() {
	wg := new(sync.WaitGroup)
	num := 0
	str := "c"
	wg.Add(1)
	go func(n int) {
		num = addOne(n)
		wg.Done()
	}(num)
	wg.Add(1)
	go func(s string) {
		str = appendStr(s)
		wg.Done()
	}(str)
	wg.Wait()
	fmt.Println(num)
	fmt.Println(str)
}
