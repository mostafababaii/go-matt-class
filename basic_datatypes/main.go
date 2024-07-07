package main

import (
	"fmt"
	"os"
)

func main() {
	a := 2
	b := 2.1
	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)

	fmt.Println(3 << 10)

	var g uint8 = 0x07
	var h = g & 0x03

	fmt.Println(h)
	fmt.Println(7 & 3)

	s := "text"
	fmt.Println(s[1:])

	var sum float64
	var n int

	for {
		var val float64
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))
}
