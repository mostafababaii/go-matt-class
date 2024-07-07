package main

import "fmt"

func do(d func() *int) {
	i := d()
	*i -= 1
	fmt.Println(*i)
}

func main() {
	//Infinite loop
	for i := 0; i < 10; i++ {
		do(func() *int {
			return &i
		})
	}
}
