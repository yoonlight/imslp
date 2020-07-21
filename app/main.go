package main

import (
	"imslp/search"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"

	//"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New() //app 선언
	//app.Settings().SetTheme(theme.LightTheme()) //밝은 테마 설정

	w := app.NewWindow("IMSLP Instrument Parsing") //window 선언
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Tchaikovsky symphony no.5")
	var data [][]string
	hbox := widget.NewHBox(
		entry, //엔트리 (문자열 입력)
		widget.NewButton("Add", func() { //종료버튼
			songURL, songName := search.Search(entry.Text)
			log.Println(songURL)
			log.Println(songName)
			data = [][]string{{songName, songURL}}
			// url = append(url, songURL)
		}),
	)
	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("IMSLP Instrument Parsing"), //레이블
			widget.NewHBox(
				entry, //엔트리 (문자열 입력)
				widget.NewButton("Add", func() { //종료버튼
					songURL, songName := search.Search(entry.Text)
					log.Println(songURL)
					log.Println(songName)
					data = [][]string{{songName, songURL}}
					widget.Refresh(hbox)
					// url = append(url, songURL)
				}),
			),
			makeTable( //makeTable함수를 사용하여 테이블 선언 및 변수 입력
				[]string{"Title", "URL"}, //칼럼(헤더)
				data,                     //데이터
			),
		),
	)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}

func makeTable(headings []string, rows [][]string) *widget.Box {

	columns := rowsToColumns(headings, rows)

	objects := make([]fyne.CanvasObject, len(columns))
	for k, col := range columns {
		box := widget.NewVBox(widget.NewLabelWithStyle(headings[k], fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		for _, val := range col {
			box.Append(widget.NewLabel(val))
		}
		objects[k] = box
	}
	return widget.NewVBox(
		fyne.NewContainerWithLayout(layout.NewGridLayout(len(columns)), objects...),
	)
}

func rowsToColumns(headings []string, rows [][]string) [][]string {
	columns := make([][]string, len(headings))
	for _, row := range rows {
		for colK := range row {
			columns[colK] = append(columns[colK], row[colK])
		}
	}
	return columns
}
