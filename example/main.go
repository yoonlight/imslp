package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	conn "imslp/connect"
	"imslp/crawler"
	"imslp/errcheck"
	imdata "imslp/imslpData"
	imparse "imslp/imslpParse"
	"imslp/input"
	"io/ioutil"
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

	for _, imslp := range lists {
		temp := strings.TrimSpace(imslp.URL) + "#tabScore2"
		log.Println(temp)
		res := conn.ConnectTLS(temp, errmsg)

		imData = crawler.IMSLPScrape(res)
		m := imparse.ParseInstr2(imData.Instr)
		imData.Instrs = m
		music = append(music, imData)
		defer res.Body.Close()
	}

	instr, _ := json.MarshalIndent(music, "", "  ")

	log.Println("Enter your json file's Title")
	fmt.Scanln(&title)
	log.Println("Complete")
	errc := ioutil.WriteFile("./"+title+".json", instr, 0)
	errcheck.CheckError(errc, "")

}
