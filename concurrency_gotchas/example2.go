package main

import "fmt"

func main() {
	ch := make(chan bool)

	go func(ok bool) {
		fmt.Println("Start")

		if ok {
			ch <- ok
		}
	}(false)

	<-ch
	fmt.Println("End")
}
