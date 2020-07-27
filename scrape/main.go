package main

import (
	"bufio"
	conn "imslp/connect"
	"imslp/crawler"
	"imslp/input"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		in     string
		errmsg = "imslp read error"
		lists  []input.List
		id     = 0
		imData []string
		music  [][]string
		title  []string
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

	for _, imslp := range lists {
		temp := strings.TrimSpace(imslp.URL) + "#tabScore2"
		log.Println(temp)
		res := conn.ConnectTLS(temp, errmsg)
		imData, title = crawler.InstrScrape(res)
		// imData = crawler.IMSLPScrape(res)
		// m := imparse.ParseInstr2(imData)
		// imData.Instrs = m
		music = append(music, imData)
		defer res.Body.Close()
	}
	log.Println(title)
	log.Println(music)
}
