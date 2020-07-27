package crawler

import (
	"imslp/errcheck"
	imdata "imslp/imslpData"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// IMSLPScrape : scrapes symphony to use infromation part crawling
func IMSLPScrape(res *http.Response) (imData imdata.IMSLPInfo) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	errcheck.CheckError(err, "imslp data Read Error")

	doc.Find("div.wi_body tr").Each(func(i int, s *goquery.Selection) {
		if s.Find("th").Text() == "Work Title\n" {
			imData.Title = s.Find("td.wi_head").Text()
		}
		if s.Find("th").Text() == "Composer\n" {
			imData.Compose = s.Find("td").Text()
		}
		if s.Find("th").Text() == "Piece Style\n" {
			imData.Style = s.Find("td").Text()
		}
		if s.Find("th").Text() == "Instrumentation\n" {
			imData.Instr = s.Find("td").Text()
		}
		if s.Find("th").Text() == "InstrDetail\n" {
			imData.Instr = s.Find("td").Text()
		}
	})
	return
}

// InstrScrape : opera 등등은 파트보를 크롤링함..
func InstrScrape(res *http.Response) (instruments []string) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	errcheck.CheckError(err, "imslp data Read Error")
	doc.Find("div#tabScore2").AddClass(".jq-ui-tabs-panel ui-widget-content ui-corner-bottom").Each(func(i int, s *goquery.Selection) {

		s.Find("a").AddClass(".eternal text").Each(func(i int, x *goquery.Selection) {
			instrument := x.Find("span").RemoveClass(".we_file_dlarrwrap").Text()
			if instrument != "" {
				instruments = append(instruments, instrument)
			}
		})
	})
	return
}
