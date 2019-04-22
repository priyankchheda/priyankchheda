// Fetch prints the content found at a URL.
// Use io.Copy instead of ioutil.ReadAll to copy the response body to
// os.Stdout without requiring a buffer large enough to hold the entire stream.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		defer resp.Body.Close()
		bytes_written, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%d\n", bytes_written)
	}
}
