package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dups3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}
	}
	fmt.Printf("%v\n", count)
}
