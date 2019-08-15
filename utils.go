package useragent

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func core(doc *html.Node, nodeParser func(*html.Node, *[]string)) []string {
	var ret []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			nodeParser(n, &ret)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return ret
}

func GetHTMLDoc(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get url")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response")
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse html code")
	}
	return doc, nil
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
