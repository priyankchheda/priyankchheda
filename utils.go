package useragent

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var userAgentStringURL = "http://www.useragentstring.com/pages/useragentstring.php"

// tagParser parses all specified tag present in HTML doc using nodeParser function
func tagParser(doc *html.Node, tag string, nodeParser func(*html.Node, *[]string)) []string {
	var ret []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			nodeParser(n, &ret)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return ret
}

// getHTMLDoc parse the html content from URL and returns html parser root Node
func getHTMLDoc(url string) (*html.Node, error) {
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

// textParser returns text of input node
func textParser(n *html.Node) string {
	var t string
	if n.Type == html.TextNode {
		t = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t = textParser(c)
	}
	return strings.TrimSpace(t)
}
