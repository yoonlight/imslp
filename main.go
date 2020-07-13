package main

import (
	conn "imslp/connect"
	"imslp/crawler"
	imslpdata "imslp/imslpData"
	imslpparse "imslp/imslpParse"
	"log"
)

func main() {
	var imslp = "https://imslp.org/wiki/Symphony_No.3%2C_Op.78_(Saint-Sa%C3%ABns%2C_Camille)"
	// var wild = "https://imslp.org/wiki/The_Wild_Dove%2C_Op.110_(Dvo%C5%99%C3%A1k%2C_Anton%C3%ADn)"
	var errmsg = ""
	res := conn.ConnectTLS(imslp, errmsg)
	title, compose, style, instrument := crawler.IMSLPScrape(res)
	log.Println(title, compose, style)
	log.Println(imslpparse.ParseInstr(instrument))
	defer res.Body.Close()
	info := imslpdata.ImportInfo(imslpparse.ParseInstr(instrument))
	log.Println(info)
	for i, arr := range info {

	}

}
