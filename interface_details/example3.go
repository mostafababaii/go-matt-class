package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Add(x, y float64) {
	p.x, p.y = p.x+x, p.y+y
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	var p Point
	p.Add(1, 2)
	fmt.Println(p)
	// Point{1, 2}.Add(2, 3) // wrong
}
