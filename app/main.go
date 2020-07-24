package main

import (
	"fmt"
	createcsv "imslp/CreateCsv"
	conn "imslp/connect"
	"imslp/crawler"
	imdata "imslp/imslpData"
	imslpparse "imslp/imslpParse"
	"imslp/input"
	"log"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	var (
		id     = 0
		data   input.List
		lists  []input.List
		infor  [][]string
		errmsg = "imslp read error"
		imData imdata.IMSLPInfo
	)
	myApp := app.New()
	myWindow := myApp.NewWindow("IMSLP")
	myWindow.Resize(fyne.NewSize(640, 480))
	vbox := widget.NewVBox()
	scroller := widget.NewVScrollContainer(vbox)
	content := widget.NewVBox(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), scroller))
	entry := widget.NewEntry()
	entry.PlaceHolder = "Tchaikovsky Symphony No.5"
	itemTitle := widget.NewLabel("Title")
	itemTitle.Alignment = fyne.TextAlignCenter

	button := widget.NewButton("Enter the Composer and Title", func() {
		if entry.Text != "" {
			data = input.Input(entry.Text, id)
			lists = append(lists, data)
			hbox := widget.NewHBox()
			on := &lists[id].IsDel
			check := widget.NewCheck("Check", func(bool) { fmt.Println("checked", on) })
			hbox.Append(widget.NewLabel(lists[id].Title))
			hbox.Append(check)
			vbox.Append(hbox)
			id++
		}
	})
	export := widget.NewButton("Export CSV", func() {
		list := make(map[int]map[string]string)
		for i, imslp := range lists {
			temp := strings.TrimSpace(imslp.URL) + "#tabScore2"
			log.Println(temp)
			res := conn.ConnectTLS(temp, errmsg)
			imData = crawler.IMSLPScrape(res)
			m := imslpparse.ParseInstr(imData.Instr)
			list[i] = m
			music := []string{imData.Title, imData.Compose, imData.Style}
			infor = append(infor, music)
			defer res.Body.Close()
		}
		title := widget.NewEntry()
		content := widget.NewForm(widget.NewFormItem("Title", title))
		dialog.ShowCustomConfirm("Enter your csv file's Title", "Submit", "Cancel", content, func(b bool) {
			log.Println("Enter your csv file's Title")
			if !b {
				log.Println("Cancle")
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
	content.Append(layout.NewSpacer())
	content.Append(export)
	content.Show()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
