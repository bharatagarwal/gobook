package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		
		readContent(data, counts)
	}

	printDuplicates(counts)
}

func readContent(data []byte, counts map[string]int) {
	fileLines := strings.Split(string(data), "\n")

	for _, line := range fileLines {
		counts[line] += 1
	}
}

func printDuplicates(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
