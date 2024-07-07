package main

import "fmt"

func do1(d func()) {
	d()
}

func main() {
	var i int

	for i = 0; i < 4; i++ {
		f := func() {
			fmt.Printf("%d -> %p\n", i, &i)
		}

		do1(f)
	}
}
