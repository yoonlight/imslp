package main

import (
	"context"
	"log"
	"net/url"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func main() {
	name := url.QueryEscape("tachikovsky symphony no.5")
	ctx := context.Background()
	url := "site:imslp.org+" + name
	log.Println(url)
	var result []googlesearch.Result
	result, _ = googlesearch.Search(ctx, url)
	log.Println(result[0].URL)
	log.Println(result[0].Title)
}
