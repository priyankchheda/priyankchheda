// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args) == 1 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err: %s", err)
				continue
			}
			defer file.Close()
			countLines(file, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(data *os.File, counts map[string]int) {
	input := bufio.NewScanner(data)
	for input.Scan() {
		counts[input.Text()]++
	}
	err := input.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "scanner err: %s", err)
	}
}
