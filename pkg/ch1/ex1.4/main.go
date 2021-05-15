package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	foundIn := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		processFiles(files, counts, foundIn)
	}

	printDuplicates(counts, foundIn)
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()] += 1
	}
}

func processFiles(files []string, counts map[string]int, foundIn map[string][]string) {

	for _, arg := range files {
		f, err := os.Open(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		countLines(f, counts, foundIn)
		f.Close()
	}
}

func printDuplicates(counts map[string]int, foundIn map[string][]string) {

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
