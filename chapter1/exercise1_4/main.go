// execise1_4 prints the names of all files in which each duplicates line
// occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileSet := make(map[string]bool)

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %s", err)
			continue
		}
		defer file.Close()

		input := bufio.NewScanner(file)
		for input.Scan() {
			line := input.Text()
			if _, ok := counts[line]; ok {
				fileSet[filename] = true
				break
			}
			counts[line]++
		}
		err = input.Err()
		if err != nil {
			fmt.Fprintf(os.Stderr, "scanner err: %s", err)
		}
	}
	for fileName, present := range fileSet {
		if present {
			fmt.Printf("%s\n", fileName)
		}
	}
}
