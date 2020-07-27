package errcheck

import "log"

// CheckError : error check with error message set by user.
func CheckError(err error, errmsg string) {
	if err != nil {
		log.Println(errmsg)
		log.Fatal(err)
	}
}
