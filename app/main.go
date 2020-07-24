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
		errmsg = "imslp read error"
		imData imdata.IMSLPInfo
		label  []*widget.Label
		checks []*widget.Check
		music  []imdata.IMSLPInfo
	)
	myApp := app.New()
	myWindow := myApp.NewWindow("IMSLP")
	myWindow.Resize(fyne.NewSize(640, 480))
	vbox := widget.NewVBox()
	// scroller := widget.NewVScrollContainer(vbox)
	// fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), scroller)
	content := widget.NewVBox()
	entry := widget.NewEntry()
	entry.PlaceHolder = "Tchaikovsky Symphony No.5"
	itemTitle := widget.NewLabel("Title")
	itemTitle.Alignment = fyne.TextAlignCenter

	button := widget.NewButton("Enter the Composer and Title", func() {
		if entry.Text != "" {
			num := id
			data = input.Input(entry.Text, num)
			lists = append(lists, data)
			hbox := widget.NewHBox()
			label = append(label, widget.NewLabel(lists[num].Title))
			check := widget.NewCheck("Check", func(on bool) {
				lists[num].IsDel = on
				fmt.Println("checked", lists[num].IsDel)
			})
			checks = append(checks, check)
			hbox.Append(label[num])
			hbox.Append(checks[num])
			vbox.Append(hbox)
			id++
		}
	})
	delete := widget.NewButton("Delete", func() {
		log.Println("delete")
		i := 0
		for i < id {
			if lists[i].IsDel == true {
				log.Println("delete", lists[i].Title)
				label[i].Hide()
				checks[i].Hide()
			}
			i++
		}
	})
	export := widget.NewButton("Export CSV", func() {
		list := make(map[int]map[string]string)
		i := 0
		for _, imslp := range lists {
			if imslp.IsDel != true {
				temp := strings.TrimSpace(imslp.URL) + "#tabScore2"
				log.Println(temp)
				res := conn.ConnectTLS(temp, errmsg)
				imData = crawler.IMSLPScrape(res)
				m := imslpparse.ParseInstr(imData.Instr)
				list[i] = m
				music = append(music, imData)
				defer res.Body.Close()
				i++
			}
		}
		title := widget.NewEntry()
		content := widget.NewForm(widget.NewFormItem("Title", title))
		dialog.ShowCustomConfirm("Enter your csv file's Title", "Submit", "Cancel", content, func(b bool) {
			log.Println("Enter your csv file's Title")
			if !b {
				log.Println("Cancle")
				return
			}
			createcsv.CreateCsv(music, list, "./"+title.Text+".csv")
			log.Println("Complete")
		}, myWindow)
	})
	content.Append(entry)
	content.Append(button)
	content.Append(itemTitle)
	content.Append(vbox)
	content.Append(layout.NewSpacer())
	content.Append(delete)
	content.Append(export)
	content.Show()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
