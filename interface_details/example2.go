package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func xyz(a int) *errFoo { return nil }

func main() {
	var err error = xyz(1)
	fmt.Println(err == nil)
	fmt.Printf("%T", err)
}
