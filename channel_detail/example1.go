package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1

	r, ok := <-ch
	fmt.Println(r, ok)

	close(ch)

	for i := 0; i < 10; i++ {
		r, ok = <-ch
		fmt.Println(r, ok)
		time.Sleep(100 * time.Millisecond)
	}
}
