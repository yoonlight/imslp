package search

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	googlesearch "github.com/rocketlaunchr/google-search"
)

// Search the IMSLP Song Link
func Search(song string) (songURL string, songName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("다시 입력하세요.", r)
		}
	}()
	var re = regexp.MustCompile(`(?m)(.*\))`)
	// song := "tachikovsky symphony no.5"
	name := url.QueryEscape(song)
	ctx := context.Background()
	url := "site:imslp.org+" + name
	// log.Println(url)
	var result []googlesearch.Result
	result, _ = googlesearch.Search(ctx, url)
	songURL = result[0].URL
	songName = re.FindAllString(result[0].Title, -1)[0]
	// log.Println(songURL)
	// log.Println(songName)
	return
}
