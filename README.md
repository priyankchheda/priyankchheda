# UserAgent
In computing, a user agent is software (a software agent) that is acting on behalf of a user, such as a web browser that "retrieves, renders and facilitates end user interaction with Web content" [(Wikipedia)](https://en.wikipedia.org/wiki/User_agent).<br/>
So basically, UserAgent identifies the browser and operating system to the web server.</br>

__Example:__ UserAgent of Google Chrome Version 76.0.3809.100 (Official Build) (64-bit)
>Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36

## Installation <a id="installation"></a>
```
go get github.com/x899/useragent
```

## Usage <a id="usage"></a>
__Source Code__
```go
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
```

__Program Output__
```
$ go run main.go

Browser Category: Crawler | Browser | Mobile Browser | Console | Offline Browser | E-mail Client | Link Checker | E-mail Collector | Validator | Feed Reader | Librarie | Cloud Platform | Other

List of browser in E-mail Client: Thunderbird

UserAgent for Thunderbird: Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.8.1.17) Gecko/20080914 Thunderbird/2.0.0.17

List of 5 Chrome UserAgent:
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36
Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/44.0.2403.155 Safari/537.36
Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.1 Safari/537.36
Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36
```

## TODO
* User Agent Analysis
* Command line tool

## Contribution
Feel Free to contribute.<br />
Please follow standard GoLang Coding Guidelines.