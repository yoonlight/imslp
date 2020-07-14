package main

import (
	"bufio"
	"encoding/csv"
	createcsv "imslp/CreateCsv"
	errcheck "imslp/ErrorCheck"
	conn "imslp/connect"
	"imslp/crawler"
	imslpparse "imslp/imslpParse"
	"os"
	"strings"
)

func main() {
	// conn.CheckHTML("https://imslp.org/wiki/Symphony_No.5%2C_Op.64_(Tchaikovsky%2C_Pyotr)")

	var url []string
	var infor [][]string

	file, _ := os.Open("./imslp.csv")

	rdr := csv.NewReader(bufio.NewReader(file))
	rdr.LazyQuotes = true

	rows, errread := rdr.ReadAll()
	errcheck.CheckError(errread, "errread")

	for _, row := range rows {
		url = append(url, row[0])
	}
	file.Close()

	var errmsg = "imslp read error"
	list := make(map[int]map[string]string)
	// log.Println(url)
	for i, imslp := range url {
		// log.Println(imslp)
		temp := strings.TrimSpace(imslp)

		res := conn.ConnectTLS(temp, errmsg)
		title, compose, style, instrument := crawler.IMSLPScrape(res)
		res.Body.Close()
		m := imslpparse.ParseInstr(instrument)
		list[i] = m
		music := []string{title, compose, style}
		infor = append(infor, music)
	}
	createcsv.CreateCsv(infor, list)

}
