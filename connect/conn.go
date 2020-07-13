package conn

import (
	"crypto/tls"
	errcheck "imslp/ErrorCheck"
	"log"
	"net/http"
)

//Connect : http connect
func Connect(url string, errmsg string) (res *http.Response) {
	res, err := http.Get(url)
	errcheck.CheckError(err, errmsg)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	log.Println(res.Status)
	return res
}

//ConnectTLS : https connect
func ConnectTLS(url string, errmsg string) (res *http.Response) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	res, err := client.Get(url)
	errcheck.CheckError(err, errmsg)
	if res.StatusCode != 200 {
		log.Panicf("status code error: %d %s", res.StatusCode, res.Status)
	}
	log.Println(res.Status)

	return res
}
