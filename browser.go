package useragent

import (
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// browserNodeParser parses each node for browser in text format
func browserNodeParser(n *html.Node, browserList *[]string) {
	*browserList = append(*browserList, textParser(n))
}

// GetBrowser returns list of valid browsers for a specified category
func GetBrowser(category string) []string {
	url := userAgentStringURL + "?typ=" + url.QueryEscape(category)
	doc, err := getHTMLDoc(url)
	if err != nil {
		log.Fatal(err)
	}
	return tagParser(doc, "h3", browserNodeParser)
}

// browserCategoryNodeParser parses each node for browser category in text format
func browserCategoryNodeParser(n *html.Node, browserCategoryList *[]string) {
	prefix := "/pages/useragentstring.php?typ="
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			if strings.HasPrefix(attr.Val, prefix) {
				category := strings.TrimPrefix(attr.Val, prefix)
				*browserCategoryList = append(*browserCategoryList, category)
			}
		}
	}
}

// GetBrowserCategory returns list of browser categories to choose from.
func GetBrowserCategory() []string {
	doc, err := getHTMLDoc(userAgentStringURL)
	if err != nil {
		log.Fatal(err)
	}
	return tagParser(doc, "a", browserCategoryNodeParser)
}
