package main

import (
	createcsv "imslp/CreateCsv"
	conn "imslp/connect"
	"imslp/crawler"
	imslpparse "imslp/imslpParse"
	"imslp/search"
	"log"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func main() {
	var (
		url    []string
		infor  [][]string
		errmsg = "imslp read error"
	)
	myApp := app.New()
	myWindow := myApp.NewWindow("IMSLP")
	myWindow.Resize(fyne.NewSize(640, 480))

	content := widget.NewVBox()
	vbox := widget.NewVBox()
	entry := widget.NewEntry()
	entry.PlaceHolder = "Tchaikovsky Symphony No.5"
	itemTitle := widget.NewLabel("Title")

	button := widget.NewButton("Enter the Composer and Title", func() {
		if entry.Text != "" {
			songURL, songName := search.Search(entry.Text)
			log.Println(songName)
			log.Println(songURL)
			url = append(url, songURL)
			vbox.Append(widget.NewLabel(songName))
		}
	})
	export := widget.NewButton("Export CSV", func() {
		list := make(map[int]map[string]string)
		for i, imslp := range url {
			temp := strings.TrimSpace(imslp) + "#tabScore2"
			log.Println(temp)
			res := conn.ConnectTLS(temp, errmsg)

			title, compose, style, instrument := crawler.IMSLPScrape(res)

			m := imslpparse.ParseInstr(instrument)
			list[i] = m
			music := []string{title, compose, style}
			infor = append(infor, music)
			defer res.Body.Close()
		}
		title := widget.NewEntry()
		content := widget.NewForm(widget.NewFormItem("Title", title))
		dialog.ShowCustomConfirm("Enter your csv file's Title", "Submit", "Cancel", content, func(b bool) {
			log.Println("Enter your csv file's Title")
			if !b {
				return
			}
			createcsv.CreateCsv(infor, list, title.Text)
			log.Println("Complete")
		}, myWindow)
	})
	content.Append(entry)
	content.Append(button)
	content.Append(itemTitle)
	content.Append(vbox)
	content.Append(export)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
