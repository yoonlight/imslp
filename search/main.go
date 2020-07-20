package main

import (
	conn "imslp/connect"
	"imslp/crawler"
	"log"
	"net/url"
)

func main() {
	name := url.QueryEscape("tachikovsky symphony no.5")

	url := "https://www.google.com/search?q=site:imslp.org+" + name
	log.Println(url)
	res := conn.ConnectTLS(url, "search error")
	log.Println(crawler.GoogleSearch(res))
	// conn.CheckHTML(url)
}
