package main

import (
	"fmt"
	"time"
)

func foo(s string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, ": ", i)
	}
}

func main() {
	go foo("1st goroutine")
	go foo("2nd goroutine")
	time.Sleep(time.Second)
	fmt.Println("Main goroutine finished")
}
