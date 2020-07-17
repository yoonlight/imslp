package createcsv

import (
	"bufio"
	"encoding/csv"
	errcheck "imslp/ErrorCheck"
	"os"

	iconv "github.com/djimenez/iconv-go"
)

// CreateCsv : csv 파일 만들어줌.
func CreateCsv(infor [][]string, list map[int]map[string]string) {
	var (
		title   string
		compose string
		style   string
	)
	file, errc := os.Create("./lumiere.csv")
	errcheck.CheckError(errc, "")
	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))
	wr.Write([]string{"Title", "Composer", "Style", "strings", "piccolo", "flute", "oboe", "english horn", "clarinet", "bass clarinet", "bassoon", "contrabassoon", "horn", "trumpet", "trombone", "tuba", "tirangle", "cymbal", "bass drum", "snare drum", "organ", "piano", "timpani", "harp"})
	for i, m := range list {
		title = infor[i][0]
		compose = infor[i][1]
		style = infor[i][2]
		result, _ := iconv.ConvertString(compose, "utf-8", "utf-8")

		wr.Write([]string{title, result, style, "1", m["piccolo"], m["flute"], m["oboe"], m["English horn"], m["clarinet"], m["bass clarinet"],
			m["bassoons"], m["contra bassoon"], m["horns"], m["trumpets"], m["trombones"], m["tuba"], m["triangle"], m["cymbals"], m["bass drum"],
			m["snare drum"], m["organ"], m["piano"], m["timpani"], m["harp"]})
		wr.Flush()
	}
}
