package conn

import (
	"crypto/tls"
	"fmt"
	"imslp/errcheck"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// CheckHTML makes html file about select html
func CheckHTML(url string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	response, err := client.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		errcheck.CheckError(err, "")

		err = ioutil.WriteFile("./html.txt", contents, 0)
		errcheck.CheckError(err, "")

		log.Println(response.Status)
	}
}
