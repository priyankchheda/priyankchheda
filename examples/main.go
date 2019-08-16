package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/x899/useragent"
)

func main() {
	browserCategories := useragent.GetBrowserCategory()
	fmt.Printf("\nBrowser Category: ")
	fmt.Println(strings.Join(browserCategories, " | "))

	category := browserCategories[5]
	browsers := useragent.GetBrowser(category)
	fmt.Printf("\nList of browser in E-mail Client: ")
	fmt.Println(strings.Join(browsers, " | "))

	browser := browsers[0]
	userAgentList := useragent.GetUserAgent(browser)

	rand.Seed(time.Now().Unix())
	fmt.Printf("\nUserAgent for %s: %s\n", browser, userAgentList[rand.Intn(len(userAgentList))])

	data := useragent.Chrome(5)
	fmt.Printf("\nList of 5 Chrome UserAgent: \n")
	for _, d := range data {
		fmt.Println(d)
	}
}
