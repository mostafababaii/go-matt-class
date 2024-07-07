package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type IntSet struct{}

func (*IntSet) String() string {
	return ""
}

func main() {
	// Declaring `m` as `map[string]Employee` is not advisable
	// Because it's not possible to take the address of a map element directly.
	// Maps in Go are dynamic structures, and the runtime may rearrange them when new elements are inserted,
	// which can invalidate previously taken addresses of map elements.
	m := map[string]*Employee{}

	m["tom"] = &Employee{Name: "Tom", Number: 1, Boss: nil, Hired: time.Now()}
	m["tom"].Number++

	m["jack"] = &Employee{Name: "Jack", Number: 3, Boss: m["tom"], Hired: time.Now()}

	fmt.Printf("%T %+[1]v\n", m["tom"])
	fmt.Printf("%T %+[1]v\n", m["jack"])

	var s IntSet

	var _ fmt.Stringer = &s
	//var _ fmt.Stringer = s // Wrong
}
