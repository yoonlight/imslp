package main

import (
	createcsv "imslp/CreateCsv"
	conn "imslp/connect"
	"imslp/crawler"
	imslpparse "imslp/imslpParse"
	"log"
)

func main() {
	var infor [][]string
	var imslp = "https://imslp.org/wiki/Symphony_No.3%2C_Op.78_(Saint-Sa%C3%ABns%2C_Camille)"
	var errmsg = ""
	res := conn.ConnectTLS(imslp, errmsg)
	title, compose, style, instrument := crawler.IMSLPScrape(res)
	res.Body.Close()
	m := imslpparse.ParseInstr(instrument)
	list := make(map[int]map[string]string)
	list[0] = m
	music := []string{title, compose, style}
	infor = append(infor, music)

	var wild = "https://imslp.org/wiki/The_Wild_Dove%2C_Op.110_(Dvo%C5%99%C3%A1k%2C_Anton%C3%ADn)"
	res = conn.ConnectTLS(wild, errmsg)
	title, compose, style, instrument = crawler.IMSLPScrape(res)
	res.Body.Close()
	m = imslpparse.ParseInstr(instrument)
	list[1] = m
	log.Println(list[0], list[1])
	music = []string{title, compose, style}
	infor = append(infor, music)

	createcsv.CreateCsv(infor, list)

}
