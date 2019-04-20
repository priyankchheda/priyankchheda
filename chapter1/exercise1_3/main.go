package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func manual() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func packageString() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	manual()
	fmt.Printf("manual function took %s\n", time.Since(start))

	start = time.Now()
	packageString()
	fmt.Printf("packageString function took %s\n", time.Since(start))
}
