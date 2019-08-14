package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHTMLReader(url string) (io.Reader, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get url")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response")
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(bodyBytes), nil
}

func text(n *html.Node) string {
	var t string
	if n.Type == html.TextNode {
		t = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t = text(c)
	}
	return t
}

func getUserAgent(htmlReader io.Reader) {
	doc, err := html.Parse(htmlReader)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, "/index.php?id=") {
						fmt.Println(text(n))
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func main() {
	browser := "BecomeBot"
	url := "http://www.useragentstring.com/pages/useragentstring.php?name="
	htmlReader, err := getHTMLReader(url + browser)
	if err != nil {
		log.Fatal(err)
	}
	getUserAgent(htmlReader)
}
