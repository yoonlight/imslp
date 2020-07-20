package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
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
		url           []string
		infor         [][]string
		errmsg        = "imslp read error"
		inputComposer string
		inputSong     string
	)

	for {
		log.Println("enter the compoeser and song:")
		fmt.Scanln(&inputComposer, &inputSong)
		n := fmt.Sprint(inputComposer + inputSong)
		if inputComposer == "exit" {
			break
		}
		songURL, songName := search.Search(n)
		log.Println(songName)
		url = append(url, songURL)

	}

	list := make(map[int]map[string]string)
	for i, imslp := range url {
		temp := strings.TrimSpace(imslp) + "#tabScore2"
		log.Println(temp)
		res := conn.ConnectTLS(temp, errmsg)

		title, compose, style, instrument := crawler.IMSLPScrape(res)
		// log.Println(crawler.InstrScrape(res), i)

		m := imslpparse.ParseInstr(instrument)
		list[i] = m
		music := []string{title, compose, style}
		infor = append(infor, music)
		defer res.Body.Close()
	}
	createcsv.CreateCsv(infor, list)
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
