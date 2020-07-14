package crawler

import (
	errcheck "imslp/ErrorCheck"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// IMSLPScrape : IMSLP 곡 정보 url을 가져오면 접속이 가능하게...
func IMSLPScrape(res *http.Response) (title string, compose string, style string, instrument string) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	errcheck.CheckError(err, "imslp data Read Error")

	doc.Find("div.wi_body tr").Each(func(i int, s *goquery.Selection) {
		// println(s.Find("th").Text())
		if s.Find("th").Text() == "Work Title\n" {
			title = s.Find("td.wi_head").Text()
			// println(title)
		}

		if s.Find("th").Text() == "Composer\n" {
			compose = s.Find("td").Text()
			// println(compose)
		}

		if s.Find("th").Text() == "Piece Style\n" {
			style = s.Find("td").Text()
			// println(style)
		}
		if s.Find("th").Text() == "Instrumentation\n" {
			instrument = s.Find("td").Text()
			// println(instrument)
		}
		if s.Find("th").Text() == "InstrDetail\n" {
			instrument = s.Find("td").Text()
			// println(instrument)
		}

		// println("---------------")

	})
	// println(title, compose, style, instrument)

	return
}
