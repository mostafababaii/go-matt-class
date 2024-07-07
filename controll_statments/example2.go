package main

import (
	"fmt"
)

func main() {
	var err error

	for {
		//Shadow declaration
		_, err := fmt.Println("Hello")
		if err != nil {
			break
		}
	}

	fmt.Println(err)
}
