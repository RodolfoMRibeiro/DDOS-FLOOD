package ddos

import (
	"fmt"
	"math/rand"
	"strconv"
)

// MULTIPLE HEADERS TO FAKE DIFFERENT REQUESTS
var acceptall = []string{
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\nAccept-Language: en-US,en;q=0.5\nAccept-Encoding: gzip, deflate",
	"Accept-Encoding: gzip, deflate",
	"Accept-Language: en-US,en;q=0.5\nAccept-Encoding: gzip, deflate",
	"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\nAccept-Language: en-US,en;q=0.5\nAccept-Charset: iso-8859-1\nAccept-Encoding: gzip",
	"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\nAccept-Charset: iso-8859-1",
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\nAccept-Charset: utf-8, iso-8859-1;q=0.5",
	"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*\nAccept-Language: en-US,en;q=0.5",
	"Accept: text/html, application/xhtml+xml, image/jxr, */*\nAccept-Encoding: gzip\nAccept-Charset: utf-8, iso-8859-1;q=0.5\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1",
	"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1\nAccept-Encoding: gzip\nAccept-Language: en-US,en;q=0.5\nAccept-Charset: utf-8, iso-8859-1;q=0.5",
	"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\nAccept-Language: en-US,en;q=0.5",
	"Accept-Charset: utf-8, iso-8859-1;q=0.5\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1",
	"Accept: text/html, application/xhtml+xml",
	"Accept-Language: en-US,en;q=0.5",
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1",
	"Accept: text/plain;q=0.8,image/png,*/*;q=0.5\nAccept-Charset: iso-8859-1",
}

// PLATFORMS, OS AND BROWSERS
var (
	platformChoices = []string{"Macintosh", "Windows", "X11"}
	macOSChoices    = []string{"Mac OS X", "Mac OS 9"}
	windowsChoices  = []string{"Windows NT 10.0", "Windows NT 6.1"}
	x11Choices      = []string{"Linux x86_64", "Linux i686"}
	browserChoices  = []string{"chrome", "ie"}
	spiderChoices   = []string{"Googlebot/2.1", "Bingbot/2.0"}
	tokenChoices    = []string{"SLCC2", "Media Center PC 6.0"}
)

// FAKE SITE REFERENCER
var (
	referers = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.fbi.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.cia.gov/index.html",
		"https://vk.com/profile.php?auto=",
		"https://www.usatoday.com/search/results?q=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
	}
)

func getUserAgent() string {
	platform := randomChoice(platformChoices)
	os := getOSForPlatform(platform)
	browser := randomChoice(browserChoices)

	switch browser {
	case "chrome":
		return buildChromeUserAgent(os)
	case "ie":
		return buildIEUserAgent(os)
	default:
		return randomChoice(spiderChoices)
	}
}

func getOSForPlatform(platform string) string {
	switch platform {
	case "Macintosh":
		return randomChoice(macOSChoices)
	case "Windows":
		return randomChoice(windowsChoices)
	case "X11":
		return randomChoice(x11Choices)
	default:
		return ""
	}
}

func buildChromeUserAgent(os string) string {
	webkit := strconv.Itoa(rand.Intn(599-500)+500) + ".0"
	version := strconv.Itoa(rand.Intn(99)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s", os, webkit, version, webkit)
}

func buildIEUserAgent(os string) string {
	version := strconv.Itoa(rand.Intn(99)) + ".0"
	engine := strconv.Itoa(rand.Intn(99)) + ".0"
	option := rand.Intn(2)

	var token string
	if option == 1 {
		token = randomChoice(tokenChoices) + "; "
	} else {
		token = ""
	}

	return fmt.Sprintf("Mozilla/5.0 (compatible; MSIE %s; %s; %sTrident/%s)", version, os, token, engine)
}

func randomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}
