package main

import (
	"fmt"
	"path/filepath"
)

type Filenamer interface {
	Filename() string
}

type Pair struct {
	Path, Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

type Fizgig struct {
	*PairWithLength
	Broken bool
}

func (p Fizgig) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d; bokrn %t", p.Path, p.Hash, p.Length, p.Broken)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 121}
	fg := Fizgig{
		&PairWithLength{Pair{"/usr/lib", "0xdead"}, 121},
		false,
	}

	fmt.Println(p)
	fmt.Println(fn)
	fmt.Println(fg)
	fmt.Println(p.Filename())
	fmt.Println(fn.Filename())
}
