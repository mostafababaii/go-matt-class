package main

import (
	"log"
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
		select {
		case m1 := <-chans[0]:
			log.Println(m1)
		case m2 := <-chans[1]:
			log.Println(m2)
		}
	}
}
