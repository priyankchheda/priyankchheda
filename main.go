package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/x899/useragent/useragent"
)

func main() {
	url := "http://www.useragentstring.com/pages/useragentstring.php"
	doc, err := useragent.GetHTMLDoc(url)
	if err != nil {
		log.Fatal(err)
	}

	browserCategory := useragent.GetBrowserCategory(doc)
	fmt.Println(browserCategory)

	browserList := useragent.GetBrowser(doc)
	fmt.Println(strings.Join(browserList, " | "))

	browser := "ABrowse"
	url = "http://www.useragentstring.com/pages/useragentstring.php?name="
	doc, err = useragent.GetHTMLDoc(url + browser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(useragent.GetUserAgent(doc))
	fmt.Println(len(useragent.GetUserAgent(doc)))

}
