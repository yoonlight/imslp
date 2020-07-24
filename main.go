package main

import (
	"bufio"
	"fmt"
	createcsv "imslp/CreateCsv"
	conn "imslp/connect"
	"imslp/crawler"
	imdata "imslp/imslpData"
	imparse "imslp/imslpParse"
	"imslp/input"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		in     string
		title  string
		errmsg = "imslp read error"
		lists  []input.List
		id     = 0
		imData imdata.IMSLPInfo
		music  []imdata.IMSLPInfo
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
		music = append(music, imData)
		defer res.Body.Close()
	}
	log.Println("Enter your csv file's Title")
	fmt.Scanln(&title)
	createcsv.CreateCsv(music, list, "./"+title+".csv")
	log.Println("Complete")
}
