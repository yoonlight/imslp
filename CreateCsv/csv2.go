package csvpkg

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"imslp/errcheck"
	imdata "imslp/imslpData"
	"os"

	"github.com/yukithm/json2csv"
)

// CreateCsv2 : csv 파일 만들어줌.
func CreateCsv2(music []imdata.IMSLPInfo, csvTitle string) {
	file, errc := os.Create("./" + csvTitle + ".csv")
	errcheck.CheckError(errc, "")

	wr := csv.NewWriter(bufio.NewWriter(file))
	// wr.Write([]string{"Title", "Composer", "Style", "strings", "piccolo", "flute", "oboe", "english horn", "clarinet", "bass clarinet", "bassoon", "contrabassoon", "horn", "trumpet", "trombone", "tuba", "tirangle", "cymbal", "bass drum", "snare drum", "organ", "piano", "timpani", "harp", "tambourine", "cannon", "bells"})
	for _, m := range music {
		instr, _ := json.Marshal(m)
		// wr.Write([]string{string(instr)})
		a, _ := json2csv.JSON2CSV(instr)
		// for _, arr := range a {
		// 	log.Println(string(arr))
		// }
		// wr.Write([]string{})

		wr.Flush()
		// jsonf.Jsonf(instr)

	}
}

// ReadCsv2 's purpose is existing data read through app's load btn.
func ReadCsv2(csvTitle string) (rows [][]string) {
	file, _ := os.Open(csvTitle)

	rdr := csv.NewReader(bufio.NewReader(file))
	rdr.LazyQuotes = true

	rows, readerr := rdr.ReadAll()
	errcheck.CheckError(readerr, "csv read err")

	file.Close()
	return
}
