package main

import (
	conn "imslp/connect"
	"imslp/crawler"
	"log"
)

func main() {
	var imslp = "https://imslp.org/wiki/Symphony_No.3%2C_Op.78_(Saint-Sa%C3%ABns%2C_Camille)"
	var errmsg = ""
	// info := new(imslpdata.IMSLPInfo)
	res := conn.ConnectTLS(imslp, errmsg)
	// title, compose, style, instrument := crawler.IMSLPScrape(res)
	log.Println(crawler.IMSLPScrape(res))
	res.Body.Close()

}
