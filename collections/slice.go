package main

import "fmt"

func main() {
	s := []byte("string")
	fmt.Println(cap(s))

	a := []int{1, 2, 3, 4}
	var b []int

	b = a

	b[3] = -1

	fmt.Println(a)

	c := []int{1, 2, 3}

	fmt.Printf("c address: %p\n", &c)
	fmt.Printf("first element address of c: %p\n", &c[0])

	c = append(c, 4)

	fmt.Printf("c address: %p\n", &c)
	fmt.Printf("first element address of c: %p\n", &c[0])
}
