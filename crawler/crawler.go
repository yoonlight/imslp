package crawler

import (
	errcheck "imslp/ErrorCheck"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// IMSLPScrape : IMSLP 곡 정보 url을 가져오면 접속이 가능하게...
func IMSLPScrape(res *http.Response) (title string, compose string, style string, instrument string) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	errcheck.CheckError(err, "imslp data Read Error")

	doc.Find("div.wi_body tr").Each(func(i int, s *goquery.Selection) {
		if s.Find("th").Text() == "Work Title\n" {
			title = s.Find("td.wi_head").Text()
		}
		if s.Find("th").Text() == "Composer\n" {
			compose = s.Find("td").Text()
		}
		if s.Find("th").Text() == "Piece Style\n" {
			style = s.Find("td").Text()
		}
		if s.Find("th").Text() == "Instrumentation\n" {
			instrument = s.Find("td").Text()
		}
		if s.Find("th").Text() == "InstrDetail\n" {
			instrument = s.Find("td").Text()
		}
	})
	return
}

// InstrScrape : opera 등등은 파트보를 크롤링함..
func InstrScrape(res *http.Response) (instrument []string) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	errcheck.CheckError(err, "imslp data Read Error")
	doc.Find("div.we_file_download plainlinks").Each(func(i int, s *goquery.Selection) {
		log.Println("sdasd")
		log.Println(s.Text())
		instrument = append(instrument, s.Find("span").Text())
		log.Println(s.Find("span").Text())
	})
	return
}
