package search

import (
	"context"
	"log"
	"net/url"
	"regexp"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func example() {
	var re = regexp.MustCompile(`(?m)(.*\))`)
	song := "tachikovsky symphony no.5"
	name := url.QueryEscape(song)
	ctx := context.Background()
	url := "site:imslp.org+" + name
	log.Println(url)
	var result []googlesearch.Result
	result, _ = googlesearch.Search(ctx, url)
	log.Println(result[0].URL)
	log.Println(re.FindAllString(result[0].Title, -1)[0])
}
