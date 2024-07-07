package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	var b [4]int

	b = a

	b[3] = -1

	fmt.Println(a)

	c := [...]byte{'a', 'b'}
	fmt.Println(c)

	if [2]int{1, 2} == [2]int{1, 2} {
		fmt.Println("YES")
	}

}
