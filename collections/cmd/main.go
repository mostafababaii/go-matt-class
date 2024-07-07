package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

type Word struct {
	Val   string
	Count int
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	var wordList []Word

	for word, count := range words {
		wordList = append(wordList, Word{Val: word, Count: count})
	}

	slices.SortFunc(wordList, func(a, b Word) int {
		return cmp.Compare(b.Count, a.Count)
	})

	fmt.Println(wordList[:3])
}
