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
		in string
		// title  string
		errmsg = "imslp read error"
		lists  []input.List
		id     = 0
		imData imdata.IMSLPInfo
		music  []imdata.IMSLPInfo
		data   interface{}
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
	for _, m := range music {

	}
	// 	for _, m := range music {
	// 	instr, _ := json.Marshal(m)
	// 	// log.Println(string(instr))
	// 	// // wr.Write([]string{string(instr)})
	// 	errj := json.Unmarshal(instr, data)
	// 	// errcheck.CheckError(errj, "asdasdsa")
	// 	a, _ := json2csv.JSON2CSV(data)
	// 	// headerStyle := headerStyleTable[c.String("header-style")]
	// 	err := printCSV(os.Stdout, a)
	// 	errcheck.CheckError(err, "print error")
	// }
	log.Println("Enter your csv file's Title")
	fmt.Scanln(&title)
	createcsv.CreateCsv(music, list, "./"+title+".csv")
	log.Println("Complete")
}
