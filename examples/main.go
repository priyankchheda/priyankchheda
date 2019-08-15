package main

import (
	"fmt"
	"strings"

	"github.com/x899/useragent"
)

func main() {
	browserCategories := useragent.GetBrowserCategory()
	fmt.Println(strings.Join(browserCategories, " | "))

	category := browserCategories[4]
	browsers := useragent.GetBrowser(category)
	fmt.Println(strings.Join(browsers, " | "))

	browser := browsers[0]
	fmt.Println(useragent.GetUserAgent(browser))
	fmt.Println(len(useragent.GetUserAgent(browser)))
}
