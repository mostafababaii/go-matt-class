package main

import "fmt"

func do1() {
	s := make([]func(), 4)
	var i int

	for i = 0; i < 4; i++ {
		i := i
		s[i] = func() {
			fmt.Printf("%d -> %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

func do2() {
	s := make([]func(), 4)
	var i int

	for i = 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d -> %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

func do3() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d -> %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

func main() {
	do1()
	fmt.Println()
	do2()
	fmt.Println()
	do3()
}
