package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2 sync.Mutex

	done := make(chan bool)

	fmt.Println("Start")

	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(time.Second)
		m2.Lock()
		defer m2.Unlock()

		fmt.Println("Signal")
		done <- true
	}()

	go func() {
		m2.Lock()
		defer m2.Unlock()
		time.Sleep(time.Second)
		m1.Lock()
		defer m1.Unlock()

		fmt.Println("Signal")
		done <- true
	}()

	<-done
	fmt.Println("Done")
	<-done
	fmt.Println("Done")
}
