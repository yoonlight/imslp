package main

import (
	"encoding/json"
	"fmt"
	pkgcsv "imslp/CreateCsv"
	conn "imslp/connect"
	"imslp/crawler"
	"imslp/errcheck"
	imdata "imslp/imslpData"
	imslpparse "imslp/imslpParse"
	"imslp/input"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
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
	load := widget.NewButton("Load don't use", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader == nil {
				return
			}
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}

			fileOpened(reader)
		}, myWindow)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		fd.Show()
		log.Println("load")
		rows := pkgcsv.ReadCsv("hello.csv")
		for _, arr := range rows {
			if arr[0] == "Title" {
				continue
			}
			num := id
			data = input.Input(arr[1]+" "+arr[0], num)
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
	export := widget.NewButton("Export CSV", func() {
		i := 0
		for _, imslp := range lists {
			if imslp.IsDel != true {
				temp := strings.TrimSpace(imslp.URL) + "#tabScore2"
				log.Println(temp)
				res := conn.ConnectTLS(temp, errmsg)
				imData = crawler.IMSLPScrape(res)
				m := imslpparse.ParseInstr2(imData.Instr)
				imData.Instrs = m
				music = append(music, imData)
				defer res.Body.Close()
				i++
			}
		}

		title := widget.NewEntry()
		content := widget.NewForm(widget.NewFormItem("Title", title))
		dialog.ShowCustomConfirm("Enter your csv file's Title", "Submit", "Cancel", content, func(b bool) {
			log.Println("Enter your json file's Title")
			if !b {
				log.Println("Cancle")
				return
			}
			data, _ := json.MarshalIndent(music, "", "  ")
			file := "./" + title.Text + ".json"
			log.Println("Complete")
			errc := ioutil.WriteFile(file, data, 0)
			errcheck.CheckError(errc, "")
			log.Println("Complete")
		}, myWindow)
	})
	content.Append(entry)
	content.Append(button)
	content.Append(itemTitle)
	content.Append(vbox)
	content.Append(layout.NewSpacer())
	content.Append(delete)
	content.Append(load)
	content.Append(export)
	content.Show()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func fileOpened(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}

	// ext := f.URI().Extension()
	// if ext == ".png" {
	// 	showImage(f)
	// } else if ext == ".txt" {
	// 	showText(f)
	// }
	err := f.Close()
	if err != nil {
		fyne.LogError("Failed to close stream", err)
	}
}
