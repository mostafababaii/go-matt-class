package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}

func producer(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	t.b = true // Unsafe at any speed
}

func consumer(vs []T, ch <-chan *T) {
	for i := range vs {
		vs[i] = *<-ch
	}
}

func main() {
	vs := make([]T, 5)
	ch := make(chan *T)

	for i := range vs {
		go producer(i, ch)
		go consumer(vs, ch)
	}

	time.Sleep(1 * time.Second) // All goroutines started
	fmt.Println(vs)
}
