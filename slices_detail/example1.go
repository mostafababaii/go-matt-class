package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	v := make([]int, 5)

	s = append(s, 1)
	v = append(v, 1)

	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", v)
}
