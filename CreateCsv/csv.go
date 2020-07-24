package createcsv

import (
	"bufio"
	"encoding/csv"
	errcheck "imslp/ErrorCheck"
	imdata "imslp/imslpData"
	"os"
)

// CreateCsv : csv 파일 만들어줌.
func CreateCsv(music []imdata.IMSLPInfo, list map[int]map[string]string, csvTitle string) {
	file, errc := os.Create(csvTitle)
	errcheck.CheckError(errc, "")

	wr := csv.NewWriter(bufio.NewWriter(file))
	wr.Write([]string{"Title", "Composer", "Style", "strings", "piccolo", "flute", "oboe", "english horn", "clarinet", "bass clarinet", "bassoon", "contrabassoon", "horn", "trumpet", "trombone", "tuba", "tirangle", "cymbal", "bass drum", "snare drum", "organ", "piano", "timpani", "harp", "tambourine", "cannon", "bells"})
	for i, m := range list {
		wr.Write([]string{music[i].Title, music[i].Compose, music[i].Style, "1", m["piccolo"], m["flute"], m["oboe"], m["English horn"], m["clarinet"], m["bass clarinet"],
			m["bassoons"], m["contra bassoon"], m["horns"], m["trumpets"], m["trombones"], m["tuba"], m["triangle"], m["cymbals"], m["bass drum"],
			m["snare drum"], m["organ"], m["piano"], m["timpani"], m["harp"], m["tambourine"], m["cannon"], m["bells"]})
		wr.Flush()
	}
}
