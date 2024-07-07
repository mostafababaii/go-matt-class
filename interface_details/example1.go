package main

import (
	"fmt"
)

type doError struct{}

func (e *doError) Error() string {
	return "Do Error"
}

func main() {
	var e error
	var se *doError

	fmt.Println(e == nil)

	e = se

	fmt.Println(e == nil)
}
