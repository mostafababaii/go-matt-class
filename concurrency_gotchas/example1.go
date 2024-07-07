package main

import (
	"fmt"
	"runtime"
	"time"
)

func do(output chan<- int, input <-chan int) {
	for output != nil {
		select {
		case m := <-input:
			output <- m
		case <-time.After(10 * time.Millisecond):
			close(output)
			output = nil
		}
	}
}

func main() {
	output := make(chan int)
	input := make(chan int)
	defer close(input)

	go do(output, input)

	input <- 1
	<-output

	fmt.Println("End")
	time.Sleep(10 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())
}
