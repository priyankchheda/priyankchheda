package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside counter")
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside /")
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

