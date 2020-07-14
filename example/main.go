package main

import (
	"bufio"
	"encoding/csv"
	errcheck "imslp/ErrorCheck"
	conn "imslp/connect"
	"imslp/crawler"
	"log"
	"os"
	"strings"
)

func main() {

	var url []string

	file, _ := os.Open("../imslp.csv")

	rdr := csv.NewReader(bufio.NewReader(file))
	rdr.LazyQuotes = true

	rows, errread := rdr.ReadAll()
	errcheck.CheckError(errread, "errread")

	for _, row := range rows {
		url = append(url, row[0])
	}
	file.Close()

	var errmsg = "imslp read error"
	for i, imslp := range url {
		temp := strings.TrimSpace(imslp) + "#tabScore2"
		log.Println(temp)
		res := conn.ConnectTLS(temp, errmsg)
		log.Println(crawler.InstrScrape(res), i)
		defer res.Body.Close()
	}
}
