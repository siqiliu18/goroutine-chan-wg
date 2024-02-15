package main

import (
	"fmt"
	channDir "proj1/channel"
	waitGroup "proj1/waitgroup"
)

// https://www.codecademy.com/resources/docs/go/goroutines
func main() {
	channDir.MyGoroutine()
	fmt.Println()
	channDir.DiffTypeChann()
	channDir.BufferChan()
	fmt.Println("------------------")
	waitGroup.Execute()
	fmt.Println()
	waitGroup.DiffTypeWg()
	fmt.Println("------------")
}
