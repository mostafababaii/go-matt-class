package main

import (
	"fmt"
	"time"
)

func flush() {
	fmt.Println("Flush")
}

func main() {
	var (
		stop = make(chan struct{}) // tells the goroutine to stop
		done = make(chan struct{}) // tells us that the goroutine exited
	)

	go func() {
		defer close(done)

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				flush()
			case <-stop:
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)

	// Elsewhere...
	close(stop) // signal the goroutine to stop
	<-done      // and wait for it to exit
}
