package main

import (
	"fmt"
	"time"
)

type X struct {
	i byte
	b bool
}

func producer1(i int, ch chan<- *X) {
	x := &X{i: byte(i)}
	ch <- x

	x.b = true // Unsafe at any speed
}

func main() {
	vs := make([]X, 5)
	ch := make(chan *X)

	for i := range vs {
		go producer1(i, ch)
	}

	time.Sleep(1 * time.Second) // All goroutines started

	// Copy quickly!
	for i := range vs {
		vs[i] = *<-ch
	}

	fmt.Println(vs)
}
