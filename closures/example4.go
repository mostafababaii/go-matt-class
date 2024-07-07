package main

import "fmt"

func do2(d func() *int) {
	i := d()
	*i -= 1
	fmt.Println(i)
}

func main() {
	var i int

	//Infinite loop
	for i = 0; i < 10; i++ {
		f := func() *int {
			fmt.Println(&i)
			return &i
		}

		do2(f)
	}
}
