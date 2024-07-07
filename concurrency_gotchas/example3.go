package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex

	done := make(chan bool)

	fmt.Println("Start")

	go func() {
		m.Lock()
		//defer m.Unlock()
	}()

	go func() {
		time.Sleep(time.Second)

		m.Lock()
		defer m.Unlock()

		fmt.Println("Signal")
		done <- true
	}()

	<-done
	fmt.Println("Done")
}
