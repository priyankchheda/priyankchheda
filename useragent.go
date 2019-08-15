package useragent

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func userAgentNodeParser(n *html.Node, userAgentList *[]string) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			if strings.HasPrefix(attr.Val, "/index.php?id=") {
				*userAgentList = append(*userAgentList, textParser(n))
			}
		}
	}
}

// GetUserAgent returns list of user agent for specified browser
func GetUserAgent(browser string) []string {
	url := userAgentStringURL + "?name="
	doc, err := GetHTMLDoc(url + browser)
	if err != nil {
		log.Fatal(err)
	}
	return tagParser(doc, "a", userAgentNodeParser)
}
