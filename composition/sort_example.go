package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (o Organs) Len() int      { return len(o) }
func (o Organs) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

type ByName struct{ Organs }
type ByWeight struct{ Organs }

func (o ByName) Less(i, j int) bool {
	return o.Organs[i].Name < o.Organs[j].Name
}

func (o ByWeight) Less(i, j int) bool {
	return o.Organs[i].Weight < o.Organs[j].Weight
}

func main() {
	organs := Organs{{"brain", 1340}, {"liver", 720}, {"heart", 290}}

	sort.Sort(ByName{organs})
	fmt.Println(organs)

	sort.Sort(ByWeight{organs})
	fmt.Println(organs)
}
