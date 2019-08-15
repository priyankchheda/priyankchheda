package useragent

import (
	"golang.org/x/net/html"
)

func browserNodeParser(n *html.Node, browserList *[]string) {
	for _, attr := range n.Attr {
		if attr.Key == "class" && attr.Val == "unterMenuName" {
			*browserList = append(*browserList, text(n))
		}
	}
}

// GetBrowser returns list of browser
func GetBrowser(doc *html.Node) []string {
	return core(doc, browserNodeParser)
}

func browserCategoryNodeParser(n *html.Node, browserCategoryList *[]string) {
	for _, attr := range n.Attr {
		if attr.Key == "class" && attr.Val == "unterMenuTitel" {
			*browserCategoryList = append(*browserCategoryList, text(n))
		}
	}
}

// GetBrowserCategory returns list of browser category
func GetBrowserCategory(doc *html.Node) []string {
	return core(doc, browserCategoryNodeParser)
}
