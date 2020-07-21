package main

import (
	"bufio"
	"encoding/csv"
	createcsv "imslp/CreateCsv"
	errcheck "imslp/ErrorCheck"
	conn "imslp/connect"
	"imslp/crawler"
	imslpparse "imslp/imslpParse"
	"imslp/search"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		url    []string
		infor  [][]string
		errmsg = "imslp read error"
		input  string
	)

	for {
		log.Println("Enter the composer and Title")
		r := bufio.NewReader(os.Stdin)
		input, _ = r.ReadString('\n')
		s := strings.TrimSpace(input)
		if s == "exit" {
			break
		}
		songURL, songName := search.Search(input)
		log.Println(songName)
		url = append(url, songURL)

	}

	list := make(map[int]map[string]string)
	for i, imslp := range url {
		temp := strings.TrimSpace(imslp) + "#tabScore2"
		log.Println(temp)
		res := conn.ConnectTLS(temp, errmsg)

		title, compose, style, instrument := crawler.IMSLPScrape(res)

		m := imslpparse.ParseInstr(instrument)
		list[i] = m
		music := []string{title, compose, style}
		infor = append(infor, music)
		defer res.Body.Close()
	}
	createcsv.CreateCsv(infor, list)
	log.Println("complete")
}

func readCSV() {
	file, _ := os.Open("./imslp.csv")

	rdr := csv.NewReader(bufio.NewReader(file))
	rdr.LazyQuotes = true

	rows, errread := rdr.ReadAll()
	log.Println(rows)
	errcheck.CheckError(errread, "errread")

	// for _, row := range rows {
	// 	url = append(url, row[0])
	// }
	file.Close()
}
