package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	loop := func(n int, ch chan<- int) {
		for {
			time.Sleep(time.Duration(n) * time.Second)
			ch <- n
		}
	}

	go loop(1, ch1)
	go loop(2, ch2)

	for i := 1; i < 12; i++ {
		fmt.Print(<-ch1, " ")
		fmt.Print(<-ch2, " ")
	}
}
