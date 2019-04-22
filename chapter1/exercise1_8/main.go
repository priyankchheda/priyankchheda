// Fetch prints the content found at a URL.
// Modify fetch to add the prefix http:// to each argument URL if it is
// missing. You might want to use strings.HasPrefix
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") &&
			!strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%s\n", body)
	}
}
