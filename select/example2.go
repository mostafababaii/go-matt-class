package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(n int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(n) * time.Second)
				ch <- n
			}
		}(i+1, chans[i])
	}

	for i := 1; i < 12; i++ {
		l1 := <-chans[0]
		l2 := <-chans[1]
		fmt.Print(l1, " ")
		fmt.Print(l2, " ")
	}
}
