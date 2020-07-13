package createcsv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	errcheck "imslp/ErrorCheck"
	"os"
)

// CreateCsv :
func CreateCsv() {
	file, errc := os.Create("./lumiere.csv")
	errcheck.CheckError(errc, "")

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// 행 읽기
	for i := range rows {
		fmt.Printf("%s ", rows[i][0])
		data := rows[i][0]

		errcheck.CheckError(errd, "")

		// csv 내용 쓰기
		wr.Write([]string{string(data), rows[i][1]})
	}
	wr.Flush()
}
