package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	createcsv "imslp/CreateCsv"
	errcheck "imslp/ErrorCheck"
	conn "imslp/connect"
	"imslp/crawler"
	imdata "imslp/imslpData"
	imparse "imslp/imslpParse"
	"imslp/input"
	"imslp/search"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		infor  [][]string
		in     string
		title  string
		errmsg = "imslp read error"
		lists  []input.List
		id     = 0
		imData imdata.IMSLPInfo
	)

	for {
		log.Println("Enter the Composer and Title")
		r := bufio.NewReader(os.Stdin)
		in, _ = r.ReadString('\n')
		s := strings.TrimSpace(in)
		if s == "exit" {
			break
		}
		lists = append(lists, input.Input(s, id))
		id++
	}

	list := make(map[int]map[string]string)
	for i, imslp := range lists {
		temp := strings.TrimSpace(imslp.URL) + "#tabScore2"
		log.Println(temp)
		res := conn.ConnectTLS(temp, errmsg)

		imData = crawler.IMSLPScrape(res)
		m := imparse.ParseInstr(imData.Instr)
		list[i] = m
		music := []string{imData.Title, imData.Compose, imData.Style}
		infor = append(infor, music)
		defer res.Body.Close()
	}
	log.Println("Enter your csv file's Title")
	fmt.Scanln(&title)
	createcsv.CreateCsv(infor, list, title)
	log.Println("Complete")
}

// existing file read
func readCSV(input string) (url []string) {
	title := "./" + input + ".csv"
	file, _ := os.Open(title)

	rdr := csv.NewReader(bufio.NewReader(file))
	rdr.LazyQuotes = true

	rows, errread := rdr.ReadAll()
	log.Println(rows)
	errcheck.CheckError(errread, "errread")

	for _, row := range rows {
		songURL, songName := search.Search(row[0])
		log.Println(songName)
		url = append(url, songURL)
	}
	file.Close()
	return
}
