package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough args")
		os.Exit(-1)
	}

	oldWord, newWord := os.Args[1], os.Args[2]
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), oldWord)
		t := strings.Join(s, newWord)

		fmt.Println(t)
	}
}
