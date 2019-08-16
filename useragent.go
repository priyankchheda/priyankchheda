package useragent

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

// userAgentNodeParser parses each node for user agent browser in text format
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
	doc, err := getHTMLDoc(url + browser)
	if err != nil {
		log.Fatal(err)
	}
	return tagParser(doc, "a", userAgentNodeParser)
}

// AOL returns list of AOL user agents
func AOL(limit int) []string {
	data := GetUserAgent("AOL")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Android returns list of Android user agents
func Android(limit int) []string {
	data := GetUserAgent("Android Webkit Browser")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// BlackBerry returns list of BlackBerry user agents
func BlackBerry(limit int) []string {
	data := GetUserAgent("BlackBerry")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// CURL returns list of curl user agents
func CURL(limit int) []string {
	data := GetUserAgent("cURL")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Chrome returns Google list of Chrome user agents
func Chrome(limit int) []string {
	data := GetUserAgent("Chrome")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Firefox returns list of Firefox user agents
func Firefox(limit int) []string {
	data := GetUserAgent("Firefox")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Iceweasel returns list of Iceweasel user agents
func Iceweasel(limit int) []string {
	data := GetUserAgent("Iceweasel")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// InternetExplorer returns Internet list of Explorer user agents
func InternetExplorer(limit int) []string {
	data := GetUserAgent("Internet Explorer")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Opera returns list of Opera user agents
func Opera(limit int) []string {
	data := GetUserAgent("Opera")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// OperaMini returns Opera list of Mini user agents
func OperaMini(limit int) []string {
	data := GetUserAgent("Opera Mini")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// OperaMobile returns Opera list of Mobile user agents
func OperaMobile(limit int) []string {
	data := GetUserAgent("Opera Mobile")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Playstation3 returns list of Playstation3 user agents
func Playstation3(limit int) []string {
	data := GetUserAgent("Playstation 3")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// PycURL returns list of python's pycurl user agents
func PycURL(limit int) []string {
	data := GetUserAgent("PycURL")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// PythonUrllib returns list of python's urllib user agents
func PythonUrllib(limit int) []string {
	data := GetUserAgent("Python-urllib")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Safari returns list of Safari user agents
func Safari(limit int) []string {
	data := GetUserAgent("Safari")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}

// Thunderbird returns list of Thunderbird user agents
func Thunderbird(limit int) []string {
	data := GetUserAgent("Thunderbird")
	if limit < len(data) {
		return data[:limit]
	}
	return data
}
