package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p point) Distance(q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := point{1, 2}
	q := point{3, 4}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)
	fmt.Printf("%T\n", point.Distance)

	p = point{3, 6}
	fmt.Println(distanceFromP(q))
}
