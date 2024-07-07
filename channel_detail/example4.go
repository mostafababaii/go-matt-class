package main

import (
	"fmt"
	"time"
)

type Z struct {
	i byte
	b bool
}

func producer2(i int, ch chan<- *Z) {
	z := &Z{i: byte(i)}
	ch <- z

	z.b = true // Unsafe at any speed
}

func main() {
	vs := make([]Z, 5)
	ch := make(chan *Z, 5)

	for i := range vs {
		go producer2(i, ch)
	}

	time.Sleep(1 * time.Second) // All goroutines started

	// Copy quickly!
	for i := range vs {
		vs[i] = *<-ch
	}

	fmt.Println(vs)
}
