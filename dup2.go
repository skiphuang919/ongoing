package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filePath := range files {
			f, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			} else {
				countLines(f, counts)
			}
			f.Close()
		}
	}
	fmt.Printf("%v\n", counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
