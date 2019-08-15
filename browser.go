package useragent

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func browserNodeParser(n *html.Node, browserList *[]string) {
	*browserList = append(*browserList, textParser(n))
}

// GetBrowser returns list of browser
func GetBrowser(category string) []string {
	url := userAgentStringURL + "?typ=Feed%20Reader"
	doc, err := GetHTMLDoc(url)
	if err != nil {
		log.Fatal(err)
	}
	return tagParser(doc, "h3", browserNodeParser)
}

// if attr.Key == "href" {
// 	if strings.HasPrefix(attr.Val, "/index.php?id=") {
// 		*userAgentList = append(*userAgentList, textParser(n))
// 	}
// }

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

// GetBrowserCategory returns list of browser category
func GetBrowserCategory() []string {
	doc, err := GetHTMLDoc(userAgentStringURL)
	if err != nil {
		log.Fatal(err)
	}
	return tagParser(doc, "a", browserCategoryNodeParser)
}
