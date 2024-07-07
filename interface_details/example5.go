package main

import (
	"fmt"
	"math"
)

type mypoint struct {
	x, y float64
}

func (p *mypoint) Distance(q mypoint) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := mypoint{1, 2}
	q := mypoint{3, 4}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)

	p = mypoint{6, 6}
	fmt.Println(distanceFromP(q))
}
