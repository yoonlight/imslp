package input

import (
	"imslp/search"
	"log"
)

// List 's purpose is easy to delete and add.
type List struct {
	ID    int
	URL   string
	Title string
	IsDel bool
}

// Input composer and title
func Input(input string, id int) (data List) {
	log.Println("Enter the Composer and Title")
	songURL, songName := search.Search(input)
	log.Println(songName)
	data = List{
		ID:    id,
		URL:   songURL,
		Title: songName,
		IsDel: false,
	}
	return
}

// History of data file read
func History() {

}
