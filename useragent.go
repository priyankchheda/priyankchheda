package useragent

import (
	"strings"

	"golang.org/x/net/html"
)

func userAgentNodeParser(n *html.Node, userAgentList *[]string) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			if strings.HasPrefix(attr.Val, "/index.php?id=") {
				*userAgentList = append(*userAgentList, text(n))
			}
		}
	}
}

// GetUserAgent returns list of user agent for specified browser
func GetUserAgent(doc *html.Node) []string {
	return core(doc, userAgentNodeParser)
}
